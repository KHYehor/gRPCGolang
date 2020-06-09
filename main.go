package main

import (
	"github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/calculate"
	"github.com/KHYehor/gRPCGolang/src/modules/server"
	"google.golang.org/grpc"
	"net"
)

var addresses = []string{
	"",
	"",
	"",
}

func getAddressListener(address string) (net.Listener, error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return lis, nil
}

func getServer(addresses []string) *grpc.Server {
	grpcServer := grpc.NewServer()
	s := &server.Server{}
	calculate.RegisterCalculateMatrixServer(grpcServer, s)
	return grpcServer
}

func main() {
	listener, err := getAddressListener("127.0.0.1:5000")
	if err != nil {
		panic("Can't listen address")
	}
	server := getServer(addresses)
	server.Serve(listener)
}
