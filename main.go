package main

import (
	"github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/calculate"
	"github.com/KHYehor/gRPCGolang/src/modules/server"
	"github.com/KHYehor/gRPCGolang/src/modules/health"
	"google.golang.org/grpc"
	"net"
)

// Rewrite to factory
func startGrpcServer(address string) (error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	s := &server.Server{}
	calculate.RegisterCalculateMatrixServer(grpcServer, s)
	grpcServer.Serve(lis)
	return nil
}

func startHealthServer(address string) (error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	s := health.
	calculate.RegisterCalculateMatrixServer(grpcServer, s)
	grpcServer.Serve(lis)
	return nil
}


func main() {
	err := startGrpcServer("127.0.0.1:5000")
	if err != nil {
		panic(err)
	}
	err = startHealthServer("127.0.0.1:6000")
	if err != nil {
		panic(err)
	}
}
