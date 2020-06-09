package server

import (
	"context"
	"errors"
	"github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/calculate"
	_ "github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/health"
	"runtime"
)

var CORES = runtime.NumCPU()
var PARALLELISM_START_NUMBER = 12

type Server struct {}

func (s *Server) MatrixSum(ctx context.Context, req *calculate.MatrixRequest) (*calculate.MatrixResponse, error) {
	if !validateMatrixEqualSize(req.Matrix1, req.Matrix2) {
		return nil, errors.New("mismatch of matrix sizes")
	}
	var result []*calculate.Array
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
