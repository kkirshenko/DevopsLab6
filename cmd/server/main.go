package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kkirshenko/DevopsLab6/internal/restaurant"
	"github.com/kkirshenko/DevopsLab6/internal/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	repo := restaurant.NewMemoryRepository()
	svc := restaurant.NewService(repo)
	eps := transport.MakeEndpoints(svc)

	mux := http.NewServeMux()
	mux.Handle("/restaurants", httptransport.NewServer(eps.Create, transport.DecodeCreateRequest, transport.EncodeResponse))
	mux.Handle("/restaurants/", httptransport.NewServer(eps.Router, transport.DecodePathRequest, transport.EncodeResponse))
	httpServer := &http.Server{Addr: ":8080", Handler: mux}

	grpcServer := grpc.NewServer()
	transport.RegisterRestaurantServer(grpcServer, transport.NewGRPCServer(eps))
	reflection.Register(grpcServer)

	errCh := make(chan error, 2)
	go func() { log.Println("HTTP listening on :8080"); errCh <- httpServer.ListenAndServe() }()
	go func() {
		l, err := net.Listen("tcp", ":9090")
		if err != nil { errCh <- err; return }
		log.Println("gRPC listening on :9090")
		errCh <- grpcServer.Serve(l)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-stop:
		log.Printf("shutdown: %s", sig)
		_ = httpServer.Shutdown(context.Background())
		grpcServer.GracefulStop()
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed { log.Fatal(err) }
	}
}
