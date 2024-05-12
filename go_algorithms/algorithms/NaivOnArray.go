package algorithms

import (
	"fmt"
	"time"
)

func NaiveOnArray(A [][]int, B [][]int) ([][]int, int64) {

	fmt.Println("Método: naiveOnArray")

	startTime := time.Now().UnixNano()
	N := len(A)
	M := len(B[0])
	P := len(B)

	if N == M && M == P {
		result := make([][]int, N)
		for i := range result {
			result[i] = make([]int, M)
		}

		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				aux := 0
				for k := 0; k < P; k++ {
					aux += A[i][k] * B[k][j]
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
