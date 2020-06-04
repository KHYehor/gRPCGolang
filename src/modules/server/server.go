package server

import (
	"context"
	"errors"
	//"github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/calculate"
	//"github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/health"
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	"github.com/KHYehor/grpcBalancer/src/grpc/health"
	"google.golang.org/grpc"
	"runtime"
	"time"
)

var CORES = runtime.NumCPU()
var PARALLELISM_START_NUMBER = 12

type Server struct {
}

func (s *Server) MatrixSum(ctx context.Context, req *calculate.MatrixRequest) (*calculate.MatrixResponse, error) {
	if !validateMatrixSumSize(req.Matrix1, req.Matrix2) {
		return nil, errors.New("mismatch of matrix sizes")
	}
	result := []*calculate.Array{}
	if len(req.Matrix1) >= PARALLELISM_START_NUMBER {
		result = calculateWithParallelism(req.GetMatrix1(), req.GetMatrix2(), matrixSum)
	} else {
		result = matrixSum(req.Matrix1, req.Matrix2)
	}
	response := &calculate.MatrixResponse{Matrix: result}
	return response, nil
}

func (s *Server) MatrixMul(ctx context.Context, req *calculate.MatrixRequest) (*calculate.MatrixResponse, error) {
	if len(req.Matrix1) > PARALLELISM_START_NUMBER {
		calculateWithParallelism(req.Matrix1, req.Matrix2, matrixMul)
	} else {
		matrixMul(req.Matrix1, req.Matrix2)
	}
	return nil, nil
}

func (s *Server) StartCheckHealth(ctx context.Context, addresses []string) {
	conn, err := grpc.Dial("")
	if err != nil {
		panic("error")
	}
	defer conn.Close()
	healthChecker := health.NewCheckHealthClient(conn)
	request := &health.HealthRequest{}
	for range time.Tick(time.Second * 1) {
		_, err := healthChecker.Health(context.Background(), request)
		if err != nil {
			// rebuild ring function
		}
	}
}
