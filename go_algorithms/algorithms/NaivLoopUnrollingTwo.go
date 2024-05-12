package algorithms

import (
	"fmt"
	"time"
)

func NaiveLoopUnrollingTwo(A [][]int, B [][]int) ([][]int, int64) {

	fmt.Println("Método: naiveLoopUnrollingTwo")

	startTime := time.Now().UnixNano()
	N := len(A)
	M := len(B[0])
	P := len(B)
	result := make([][]int, N)
	for i := range result {
		result[i] = make([]int, M)
	}

	if N == M && M == P {
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				aux := 0
				if P%2 == 0 {
					for k := 0; k < P; k += 2 {
						aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j]
					}
				} else {
					PP := P - 1
					for k := 0; k < PP; k += 2 {
						aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j]
					}
					aux += A[i][PP] * B[PP][j]
				}
				result[i][j] = aux
			}
		}

		endTime := time.Now().UnixNano()
		executionTime := endTime - startTime
		return result, executionTime
	} else {
		fmt.Println("Las dimensiones de las matrices no son iguales.")
		return nil, 0
	}
}
