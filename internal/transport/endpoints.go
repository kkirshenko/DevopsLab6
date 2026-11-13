package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/kkirshenko/DevopsLab6/internal/restaurant"
)

type Endpoints struct {
	Create endpoint.Endpoint
	Get endpoint.Endpoint
	ChangeRating endpoint.Endpoint
	Router endpoint.Endpoint
}
type createRequest struct { Restaurant restaurant.Restaurant }
type idRequest struct { ID string }
type ratingRequest struct { ID string; Rating float32 }

func MakeEndpoints(s restaurant.Service) Endpoints {
	create := endpoint.Endpoint(func(ctx context.Context, request interface{}) (interface{}, error) {
		r := request.(createRequest)
		return s.Create(ctx, r.Restaurant)
	})
	get := endpoint.Endpoint(func(ctx context.Context, request interface{}) (interface{}, error) {
		r := request.(idRequest)
		return s.GetByID(ctx, r.ID)
	})
	change := endpoint.Endpoint(func(ctx context.Context, request interface{}) (interface{}, error) {
		r := request.(ratingRequest)
		return struct{}{}, s.ChangeRating(ctx, r.ID, r.Rating)
	})
	router := endpoint.Endpoint(func(ctx context.Context, request interface{}) (interface{}, error) {
		switch r := request.(type) {
		case idRequest: return get(ctx, r)
		case ratingRequest: return change(ctx, r)
		default: return nil, restaurant.ErrInvalid
		}
	})
	return Endpoints{Create:create, Get:get, ChangeRating:change, Router:router}
}
