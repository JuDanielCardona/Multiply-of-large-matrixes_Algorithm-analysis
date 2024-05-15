package algorithms

import (
	"fmt"
	"time"
)

func NaivLoopUnrollingFour(A [][]int, B [][]int) ([][]int, int64) {

	fmt.Println("Método: naiveLoopUnrollingFour")

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
				if P%4 == 0 {
					for k := 0; k < P; k += 4 {
						aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
					}
				} else if P%4 == 1 {
					PP := P - 1
					for k := 0; k < PP; k += 4 {
						aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
					}
					aux += A[i][PP] * B[PP][j]
				} else if P%4 == 2 {
					PP := P - 2
					PPP := P - 1
					for k := 0; k < PP; k += 4 {
						aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
					}
					aux += A[i][PP]*B[PP][j] + A[i][PPP]*B[PPP][j]
				} else {
					PP := P - 3
					PPP := P - 2
					PPPP := P - 1
					for k := 0; k < PP; k += 4 {
						aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
					}
					aux += A[i][PP]*B[PP][j] + A[i][PPP]*B[PPP][j] + A[i][PPPP]*B[PPPP][j]
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
