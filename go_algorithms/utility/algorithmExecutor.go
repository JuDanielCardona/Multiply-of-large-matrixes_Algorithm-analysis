package utility

import (
	"fmt"
	"go_algorithms/algorithms"
)

func Init_GoExecution(matrixA [][]int, matrixB [][]int) {
	executionTimes := make(map[string]int64)

	result, time := algorithms.NaiveOnArray(matrixA, matrixB)
	executionTimes["NaiveOnArray"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.NaiveLoopUnrollingTwo(matrixA, matrixB)
	executionTimes["NaiveLoopUnrollingTwo"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.NaivLoopUnrollingFour(matrixA, matrixB)
	executionTimes["NaiveLoopUnrollingFour"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.WinogradOriginal(matrixA, matrixB)
	executionTimes["WinogradOriginal"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.WinogradScaled(matrixA, matrixB)
	executionTimes["WinogradScaled"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.StrassenNaiv(matrixA, matrixB)
	executionTimes["StrassenNaiv"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.StrassenWinograd(matrixA, matrixB)
	executionTimes["StrassenWinograd"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.RowByColumnSequential(matrixA, matrixB)
	executionTimes["RowByColumnSequential"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.RowByColumnParallel(matrixA, matrixB)
	executionTimes["RowByColumnParallel"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.RowByColumnEnhanced(matrixA, matrixB)
	executionTimes["RowByColumnEnhanced"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.RowByRowSequential(matrixA, matrixB)
	executionTimes["RowByRowSequential"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.RowByRowParallel(matrixA, matrixB)
	executionTimes["RowByRowParallel"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.RowByColumnEnhanced(matrixA, matrixB)
	executionTimes["RowByColumnEnhanced"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.ColumnByColumnSequential(matrixA, matrixB)
	executionTimes["ColumnByColumnSequential"] = time
	fmt.Print(time, result, "\n\n")

	result, time = algorithms.ColumnByColumnParallel(matrixA, matrixB)
	executionTimes["ColumnByColumnParallel"] = time
	fmt.Print(time, result, "\n\n")
}
