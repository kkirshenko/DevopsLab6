package transport

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

type grpcServer struct {
	create *kitgrpc.Server
	get *kitgrpc.Server
	change *kitgrpc.Server
}
func NewGRPCServer(e Endpoints) *grpcServer {
	decode := func(_ context.Context, req interface{}) (interface{}, error) {
		m := req.(*structpb.Struct).AsMap()
		switch m["method"] {
		case "Create":
			return createRequest{Restaurant: restaurantFromMap(m)}, nil
		case "GetByID":
			return idRequest{ID:stringValue(m["id"])}, nil
		case "ChangeRating":
			return ratingRequest{ID:stringValue(m["id"]), Rating:float32(numberValue(m["rating"]))}, nil
		default: return nil, status.Error(codes.InvalidArgument, "unknown method")
		}
	}
	encode := func(_ context.Context, response interface{}) (interface{}, error) {
		switch v := response.(type) {
		case string: return structpb.NewStruct(map[string]interface{}{"id":v})
		case restaurant.Restaurant:
			return structpb.NewStruct(map[string]interface{}{"id":v.ID,"name":v.Name,"address":v.Address,"rating":float64(v.Rating),"open":v.Open})
		default: return structpb.NewStruct(map[string]interface{}{"ok":true})
		}
	}
	return &grpcServer{
		create:kitgrpc.NewServer(e.Create, decode, encode),
		get:kitgrpc.NewServer(e.Router, decode, encode),
		change:kitgrpc.NewServer(e.Router, decode, encode),
	}
}
func (s *grpcServer) Create(ctx context.Context, req *structpb.Struct) (*structpb.Struct,error) { _, r, err := s.create.ServeGRPC(ctx,req); if err != nil{return nil,err}; return r.(*structpb.Struct),nil }
func (s *grpcServer) GetByID(ctx context.Context, req *structpb.Struct) (*structpb.Struct,error) { _, r, err := s.get.ServeGRPC(ctx,req); if err != nil{return nil,err}; return r.(*structpb.Struct),nil }
func (s *grpcServer) ChangeRating(ctx context.Context, req *structpb.Struct) (*structpb.Struct,error) { _, r, err := s.change.ServeGRPC(ctx,req); if err != nil{return nil,err}; return r.(*structpb.Struct),nil }

func RegisterRestaurantServer(reg grpc.ServiceRegistrar, srv *grpcServer) {
	reg.RegisterService(&grpc.ServiceDesc{ServiceName:"restaurant.v1.RestaurantService", HandlerType:(*restaurantServiceServer)(nil), Methods:[]grpc.MethodDesc{
		{MethodName:"Create",HandlerFunc:handlerCreate},
		{MethodName:"GetByID",HandlerFunc:handlerGet},
		{MethodName:"ChangeRating",HandlerFunc:handlerChange},
	}}, srv)
}
type restaurantServiceServer interface { Create(context.Context,*structpb.Struct)(*structpb.Struct,error); GetByID(context.Context,*structpb.Struct)(*structpb.Struct,error); ChangeRating(context.Context,*structpb.Struct)(*structpb.Struct,error) }
func handlerCreate(s interface{},ctx context.Context,dec func(interface{})error,i grpc.UnaryServerInterceptor)(interface{},error){ in:=new(structpb.Struct); if e:=dec(in);e!=nil{return nil,e}; call:=func(c context.Context,r interface{})(interface{},error){return s.(restaurantServiceServer).Create(c,r.(*structpb.Struct))}; if i==nil{return call(ctx,in)}; return i(ctx,in,&grpc.UnaryServerInfo{Server:s,FullMethod:"/restaurant.v1.RestaurantService/Create"},call)}
func handlerGet(s interface{},ctx context.Context,dec func(interface{})error,i grpc.UnaryServerInterceptor)(interface{},error){ in:=new(structpb.Struct); if e:=dec(in);e!=nil{return nil,e}; call:=func(c context.Context,r interface{})(interface{},error){return s.(restaurantServiceServer).GetByID(c,r.(*structpb.Struct))}; if i==nil{return call(ctx,in)}; return i(ctx,in,&grpc.UnaryServerInfo{Server:s,FullMethod:"/restaurant.v1.RestaurantService/GetByID"},call)}
func handlerChange(s interface{},ctx context.Context,dec func(interface{})error,i grpc.UnaryServerInterceptor)(interface{},error){ in:=new(structpb.Struct); if e:=dec(in);e!=nil{return nil,e}; call:=func(c context.Context,r interface{})(interface{},error){return s.(restaurantServiceServer).ChangeRating(c,r.(*structpb.Struct))}; if i==nil{return call(ctx,in)}; return i(ctx,in,&grpc.UnaryServerInfo{Server:s,FullMethod:"/restaurant.v1.RestaurantService/ChangeRating"},call)}
func stringValue(v interface{}) string { s,_:=v.(string); return s }
func numberValue(v interface{}) float64 { n,_:=v.(float64); return n }
func restaurantFromMap(m map[string]interface{}) restaurant.Restaurant { return restaurant.Restaurant{ID:stringValue(m["id"]),Name:stringValue(m["name"]),Address:stringValue(m["address"]),Rating:float32(numberValue(m["rating"])),Open:m["open"]==true} }
