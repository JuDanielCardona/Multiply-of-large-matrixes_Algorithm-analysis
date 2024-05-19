package utility

import (
	"fmt"
	"go_algorithms/algorithms"
	"go_algorithms/database"
)

func Init_execution(matrixA [][]int, matrixB [][]int, test string) {
	isPrint := true

	result, time := algorithms.NaiveOnArray(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "NaiveOnArray", time)
	print(time, result, isPrint)

	result, time = algorithms.NaiveLoopUnrollingTwo(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "NaiveLoopUnrollingTwo", time)
	print(time, result, isPrint)

	result, time = algorithms.NaivLoopUnrollingFour(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "NaivLoopUnrollingFour", time)
	print(time, result, isPrint)

	result, time = algorithms.WinogradOriginal(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "WinogradOriginal", time)
	print(time, result, isPrint)

	result, time = algorithms.WinogradScaled(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "WinogradScaled", time)
	print(time, result, isPrint)

	result, time = algorithms.StrassenNaiv(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "StrassenNaiv", time)
	print(time, result, isPrint)

	result, time = algorithms.StrassenWinograd(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "StrassenWinograd", time)
	print(time, result, isPrint)

	result, time = algorithms.RowByColumnSequential(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "RowByColumnSequential", time)
	print(time, result, isPrint)

	result, time = algorithms.RowByColumnParallel(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "RowByColumnParallel", time)
	print(time, result, isPrint)

	result, time = algorithms.RowByColumnEnhanced(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "RowByColumnEnhanced", time)
	print(time, result, isPrint)

	result, time = algorithms.RowByRowSequential(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "RowByRowSequential", time)
	print(time, result, isPrint)

	result, time = algorithms.RowByRowParallel(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "RowByRowParallel", time)
	print(time, result, isPrint)

	result, time = algorithms.RowByColumnEnhanced(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "RowByColumnEnhanced", time)
	print(time, result, isPrint)

	result, time = algorithms.ColumnByColumnSequential(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "ColumnByColumnSequential", time)
	print(time, result, isPrint)

	result, time = algorithms.ColumnByColumnParallel(matrixA, matrixB)
	database.RegisterInfo(test, "Golang", "ColumnByColumnParallel", time)
	print(time, result, isPrint)

	fmt.Print("All algorithms executed correctly\n\n")
}

func print(time int64, matrix [][]int, print bool) {
	if print {
		for i := 0; i < 5 && i < len(matrix); i++ {
			fmt.Printf("%d ", matrix[i][0])
		}
	} else {
		totalValues := len(matrix) * len(matrix[0])
		fmt.Printf(" + %d other values\n = %dns", totalValues-5, time)
	}
}
