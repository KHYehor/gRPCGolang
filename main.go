package main

import (
	"context"
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	"github.com/KHYehor/gRPCBalancer/src/server"
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

func getLoadBalancer(addresses []string) *grpc.Server {
	grpcServer := grpc.NewServer()
	s := &server.LoadBalancer{}
	s.InitServers(context.Background(), addresses)
	calculate.RegisterCalculateMatrixServer(grpcServer, s)
	return grpcServer
}

func main() {
	listener, err := getAddressListener("127.0.0.1:5000")
	if err != nil {
		panic("Can't listen address")
	}
	server := getLoadBalancer(addresses)
	server.Serve(listener)
}
