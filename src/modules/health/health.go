package health

import (
	"context"
	"github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/health"
)

type HealthServer struct {}

func (hs *HealthServer) Health(ctx context.Context, req *health.HealthRequest) (*health.HealthResponse, error) {
	// Free memory in MB
	//si := sysinfo.Get().FreeHighRam / 1024
	return &health.HealthResponse{MemoryAllocated: /* si */ 12}, nil
}
