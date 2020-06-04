package server

import (
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	"sync"
)

type matrixOperation func([]*calculate.Array, []*calculate.Array) []*calculate.Array


func matrixMul(matrix1 []*calculate.Array, matrix2 []*calculate.Array) []*calculate.Array {
	return []*calculate.Array{}
}

func validateMatrixSumSize(matrix1 []*calculate.Array, matrix2 []*calculate.Array) bool {
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

func matrixSum(matrix1 []*calculate.Array, matrix2 []*calculate.Array) []*calculate.Array {
	// Init result array
	var result = []*calculate.Array{}
	for i := 0; i < len(matrix1); i++ {
		result = append(result, &calculate.Array{})
	}
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
	// Copy appropirate first array
	for i := from; i < to; i++ {
		copiedMatrix1 = append(copiedMatrix1, matrix1[i])
		copiedMatrix2 = append(copiedMatrix2, matrix2[i])
	}
	//// Copy appropirate second array
	//for i := 0; i < len(matrix2); i++ {
	//	subArray := matrix2[i]
	//	subArray.Digit = []float64{}
	//	for j := from; j < to; j++ {
	//		subArray.Digit = append(subArray.Digit, matrix2[i].Digit[j])
	//	}
	//	copiedMatrix2 = append(copiedMatrix2, subArray)
	//}
	return copiedMatrix1, copiedMatrix2
}

func calculateWithParallelism(matrix1 []*calculate.Array, matrix2 []*calculate.Array, operation matrixOperation) []*calculate.Array {
	var wg sync.WaitGroup
	wg.Add(CORES)
	step := len(matrix1)/CORES
	var result []*calculate.Array
	var temp [][]*calculate.Array
	// Init array of necessary size
	for i := 0; i < CORES; i++ {
		temp = append(temp, nil)
	}
	for coresCounter := 0; coresCounter < CORES; coresCounter++ {
		go func(i int) {
			defer wg.Done()
			copiedMatrix1, copiedMatrix2 := copyMatrixes(matrix1, matrix2, step * i, step * (i + 1))
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