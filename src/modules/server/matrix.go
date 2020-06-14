package server

import (
	"github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/calculate"
	"sync"
)

type matrixOperation func([]*calculate.Array, []*calculate.Array) []*calculate.Array

// Will be done soon
func matrixMul(matrix1 []*calculate.Array, matrix2 []*calculate.Array) []*calculate.Array {
	return []*calculate.Array{}
}

func validateMatrixEqualSize(matrix1 []*calculate.Array, matrix2 []*calculate.Array) bool {
	if len(matrix1) != len(matrix2) {
		return false
	}
	for i := 0; i < len(matrix1); i++ {
		if len(matrix1[i].Digit) != len(matrix2[i].Digit) {
			return false
		}
	}
	return true
}

func createEmptyArray(size int, row int) []*calculate.Array {
	var empty []*calculate.Array
	for i := 0; i < size; i++ {
		emptyRow := []float64{}
		for j := 0; j < row; j++ {
			emptyRow = append(emptyRow, 0)
		}
		empty = append(empty, &calculate.Array{Digit: emptyRow })
	}
	return empty
}

func matrixSum(matrix1 []*calculate.Array, matrix2 []*calculate.Array) []*calculate.Array {
	// Init result array
	var result = createEmptyArray(len(matrix1), len(matrix1[0].Digit))
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1); j++ {
			result[i].Digit[j] = matrix1[i].Digit[j] + matrix2[i].Digit[j]
		}
	}
	return result
}

func copyMatrixes(matrix1 []*calculate.Array, matrix2 []*calculate.Array, from int, to int) ([]*calculate.Array, []*calculate.Array) {
	var copiedMatrix1 []*calculate.Array
	var copiedMatrix2 []*calculate.Array
	for i := from; i < to; i++ {
		copiedMatrix1 = append(copiedMatrix1, matrix1[i])
		copiedMatrix2 = append(copiedMatrix2, matrix2[i])
	}
	return copiedMatrix1, copiedMatrix2
}

func calculateWithParallelism(matrix1 []*calculate.Array, matrix2 []*calculate.Array, operation matrixOperation) []*calculate.Array {
	var wg sync.WaitGroup
	wg.Add(CORES)
	step := len(matrix1) / CORES
	var result []*calculate.Array
	var temp [][]*calculate.Array
	// Init array of necessary size
	for i := 0; i < CORES; i++ {
		temp = append(temp, nil)
	}
	for coresCounter := 0; coresCounter < CORES; coresCounter++ {
		go func(i int) {
			defer wg.Done()
			copiedMatrix1, copiedMatrix2 := copyMatrixes(matrix1, matrix2, step*i, step*(i+1))
			temp[i] = operation(copiedMatrix1, copiedMatrix2)
		}(coresCounter)
	}
	wg.Wait()
	// Concating back
	for _, value := range temp {
		for _, value := range value {
			result = append(result, value)
		}
	}
	return result
}
