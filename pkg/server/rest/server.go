package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	ps "github.com/f-go/petstore/pkg/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// ...
func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := ps.RegisterPetstoreServiceHandlerFromEndpoint(ctx, mux, ":" + grpcPort, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}
	if err := ps.RegisterMaintenanceServiceHandlerFromEndpoint(ctx, mux, ":" + grpcPort, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down HTTP/TREST server...")
		}
		_, cancel := context.WithTimeout(ctx, 5 * time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
