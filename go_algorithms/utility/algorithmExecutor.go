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

	result, time = algorithms.NaiveLoopUnrollingFour(matrixA, matrixB)
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

	result, time = algorithms.StrassenNaiv(matrixA, matrixB)
	executionTimes["StrassenNaiv"] = time
	fmt.Print(time, result, "\n\n")

}
