package algorithms

import (
	"fmt"
	"time"
)

func WinogradScaled(A, B [][]int) ([][]int, int64) {

	fmt.Println("Método: WinogradScaled")

	startTime := time.Now().UnixNano()
	m := len(A)
	n := len(B[0])
	c := make([][]int, m)
	d := make([][]int, m)

	// Inicializar las matrices c y d
	for i := range c {
		c[i] = make([]int, n)
		d[i] = make([]int, n)
	}

	// Calcular d
	for i := 0; i < m; i++ {
		for k := 0; k < n/2; k++ {
			d[i][2*k] = A[i][2*k+1] + B[2*k][i]
			d[i][2*k+1] = A[i][2*k] + B[2*k+1][i]
		}
	}

	// Calcular c
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := -d[i][0] * d[i][1]
			for k := 1; k < n/2; k++ {
				sum += d[i][2*k] * d[i][2*k+1]
			}
			c[i][j] = sum
		}
	}

	// Calcular el producto final
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		for j := range result[i] {
			result[i][j] = c[i][j]
			for k := 0; k < n/2; k++ {
				result[i][j] += (A[i][2*k] + B[2*k+1][j]) * (A[i][2*k+1] + B[2*k][j])
			}
		}
	}

	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime
	return result, executionTime
}
