package algorithms

import (
	"fmt"
	"time"
)

func WinogradOriginal(A [][]int, B [][]int) ([][]int, int64) {
	fmt.Println("Método: WinogradOriginal")

	startTime := time.Now().UnixNano()

	N := len(B[0])
	P := len(B)
	M := len(A)

	Result := make([][]int, M)
	for i := range Result {
		Result[i] = make([]int, N)
	}

	var i, j, k int
	var aux int
	upsilon := P % 2
	gamma := P - upsilon
	y := make([]int, M)
	z := make([]int, N)
	for i = 0; i < M; i++ {
		aux = 0
		for j = 0; j < gamma; j += 2 {
			aux += A[i][j] * A[i][j+1]
		}
		y[i] = aux
	}
	for i = 0; i < N; i++ {
		aux = 0
		for j = 0; j < gamma; j += 2 {
			aux += B[j][i] * B[j+1][i]
		}
		z[i] = aux
	}

	if upsilon == 1 {
		PP := P - 1
		for i = 0; i < M; i++ {
			for k = 0; k < N; k++ {
				aux = 0
				for j = 0; j < gamma; j += 2 {
					aux += (A[i][j] + B[j+1][k]) * (A[i][j+1] + B[j][k])
				}
				Result[i][k] = aux - y[i] - z[k] + A[i][PP]*B[PP][k]
			}
		}
	} else {
		for i = 0; i < M; i++ {
			for k = 0; k < N; k++ {
				aux = 0
				for j = 0; j < gamma; j += 2 {
					aux += (A[i][j] + B[j+1][k]) * (A[i][j+1] + B[j][k])
				}
				Result[i][k] = aux - y[i] - z[k]
			}
		}
	}

	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime
	return Result, executionTime
}
