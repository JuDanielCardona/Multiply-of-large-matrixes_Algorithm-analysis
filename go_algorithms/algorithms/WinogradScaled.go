package algorithms

import (
	"fmt"
	"math"
	"time"
)

func WinogradScaled(A, B [][]int) ([][]int, int64) {

	fmt.Println("Método: WinogradScaled")

	startTime := time.Now().UnixNano()

	N := len(B[0])
	P := len(B)
	M := len(A)

	CopyA := make([][]int, N)
	for i := range CopyA {
		CopyA[i] = make([]int, P)
		copy(CopyA[i], A[i])
	}

	CopyB := make([][]int, P)
	for i := range CopyB {
		CopyB[i] = make([]int, M)
		copy(CopyB[i], B[i])
	}

	// Scaling factors
	a := NormInf(A, N, P)
	b := NormInf(B, P, M)
	lambda := math.Floor(0.5 + math.Log(float64(b)/float64(a))/math.Log(4))

	// Scaling
	MultiplyWithScalar(A, CopyA, N, P, math.Pow(2, float64(lambda)))
	MultiplyWithScalar(B, CopyB, P, M, math.Pow(2, -float64(lambda)))

	// Using Winograd with scaled matrices
	result := Winograd(CopyA, CopyB)

	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime
	return result, executionTime
}

func NormInf(matrix [][]int, rows, cols int) int {
	maxNorm := 0
	for i := 0; i < rows; i++ {
		rowSum := 0
		for j := 0; j < cols; j++ {
			rowSum += int(math.Abs(float64(matrix[i][j])))
		}
		if rowSum > maxNorm {
			maxNorm = rowSum
		}
	}
	return maxNorm
}

func MultiplyWithScalar(A, B [][]int, rows, cols int, scalar float64) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			B[i][j] = int(float64(A[i][j]) * scalar)
		}
	}
}

func Winograd(A [][]int, B [][]int) [][]int {

	N := len(B[0])
	P := len(B)
	M := len(A)

	Result := make([][]int, M)
	for i := range Result {
		Result[i] = make([]int, N)
	}

	var aux int
	upsilon := P % 2
	gamma := P - upsilon
	y := make([]int, M)
	z := make([]int, N)
	for i := 0; i < M; i++ {
		aux = 0
		for j := 0; j < gamma; j += 2 {
			aux += A[i][j] * A[i][j+1]
		}
		y[i] = aux
	}
	for i := 0; i < N; i++ {
		aux = 0
		for j := 0; j < gamma; j += 2 {
			aux += B[j][i] * B[j+1][i]
		}
		z[i] = aux
	}

	if upsilon == 1 {
		PP := P - 1
		for i := 0; i < M; i++ {
			for k := 0; k < N; k++ {
				aux = 0
				for j := 0; j < gamma; j += 2 {
					aux += (A[i][j] + B[j+1][k]) * (A[i][j+1] + B[j][k])
				}
				Result[i][k] = aux - y[i] - z[k] + A[i][PP]*B[PP][k]
			}
		}
	} else {
		for i := 0; i < M; i++ {
			for k := 0; k < N; k++ {
				aux = 0
				for j := 0; j < gamma; j += 2 {
					aux += (A[i][j] + B[j+1][k]) * (A[i][j+1] + B[j][k])
				}
				Result[i][k] = aux - y[i] - z[k]
			}
		}
	}

	return Result
}
