package gha

import (
	// "crypto/tls"
	// "crypto/x509"
	"context"

	"github.com/f-go/petstore/pkg/server/grpc"
	"github.com/f-go/petstore/pkg/server/rest"
	petstore "github.com/f-go/petstore/pkg/service"
)

// ...
func RunServer() error {
	ctx := context.Background()
	petstoreService := petstore.NewPetstoreService()
	maintenanceService := petstore.NewMaintenanceService()

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, "9090", "8080")
	}()
	return grpc.RunServer(ctx, petstoreService, maintenanceService, "9090")
}
