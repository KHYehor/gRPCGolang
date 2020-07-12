package main

import (
	"errors"
	"github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/calculate"
	healthgrpc "github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/health"
	"github.com/KHYehor/gRPCGolang/src/modules/health"
	"github.com/KHYehor/gRPCGolang/src/modules/server"
	"google.golang.org/grpc"
	"net"
)

// Factory of gRPC servers
func grpcServerFactory(serverType string, address string) (error) {
	grpcServer := grpc.NewServer()
	if serverType == "calculate" {
		calculate.RegisterCalculateMatrixServer(grpcServer, &server.Server{})
	} else if serverType == "health" {
		healthgrpc.RegisterCheckHealthServer(grpcServer, &health.HealthServer{})
	} else {
		return errors.New("unknown server type")
	}
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer.Serve(lis)
	return nil
}

func main() {
	err := grpcServerFactory("calculate", "0.0.0.0:5000")
	if err != nil {
		panic(err)
	}
	err = grpcServerFactory("health", "0.0.0.0:6000")
	if err != nil {
		panic(err)
	}
}
