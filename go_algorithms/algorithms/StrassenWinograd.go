package algorithms

import (
	"fmt"
	"time"
)

func StrassenWinograd(matrixA, matrixB [][]int) ([][]int, int64) {

	fmt.Println("Método: strassenWinograd")

	startTime := time.Now().UnixNano()
	result := multiplication_SW(matrixA, matrixB)
	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime
	return result, executionTime
}

func sumMatrix(a, b [][]int) [][]int {
	n := len(a)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

func subtractMatrix(a, b [][]int) [][]int {
	n := len(a)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = a[i][j] - b[i][j]
		}
	}
	return result
}

func multiplication_SW(a, b [][]int) [][]int {
	n := len(a)

	if n == 1 {
		result := make([][]int, 1)
		result[0] = make([]int, 1)
		result[0][0] = a[0][0] * b[0][0]
		return result
	}

	newSize := n / 2

	a11 := make([][]int, newSize)
	a12 := make([][]int, newSize)
	a21 := make([][]int, newSize)
	a22 := make([][]int, newSize)
	b11 := make([][]int, newSize)
	b12 := make([][]int, newSize)
	b21 := make([][]int, newSize)
	b22 := make([][]int, newSize)
	for i := 0; i < newSize; i++ {
		a11[i] = make([]int, newSize)
		a12[i] = make([]int, newSize)
		a21[i] = make([]int, newSize)
		a22[i] = make([]int, newSize)
		b11[i] = make([]int, newSize)
		b12[i] = make([]int, newSize)
		b21[i] = make([]int, newSize)
		b22[i] = make([]int, newSize)
	}

	// Divide matrices into submatrices
	for i := 0; i < newSize; i++ {
		for j := 0; j < newSize; j++ {
			a11[i][j] = a[i][j]
			a12[i][j] = a[i][j+newSize]
			a21[i][j] = a[i+newSize][j]
			a22[i][j] = a[i+newSize][j+newSize]

			b11[i][j] = b[i][j]
			b12[i][j] = b[i][j+newSize]
			b21[i][j] = b[i+newSize][j]
			b22[i][j] = b[i+newSize][j+newSize]
		}
	}

	// Calculate submatrices
	m1 := multiplication_SW(sumMatrix(a11, a22), sumMatrix(b11, b22))
	m2 := multiplication_SW(sumMatrix(a21, a22), b11)
	m3 := multiplication_SW(a11, subtractMatrix(b12, b22))
	m4 := multiplication_SW(a22, subtractMatrix(b21, b11))
	m5 := multiplication_SW(sumMatrix(a11, a12), b22)
	m6 := multiplication_SW(subtractMatrix(a21, a11), sumMatrix(b11, b12))
	m7 := multiplication_SW(subtractMatrix(a12, a22), sumMatrix(b21, b22))

	// Calculate result submatrices
	c11 := sumMatrix(subtractMatrix(sumMatrix(m1, m4), m5), m7)
	c12 := sumMatrix(m3, m5)
	c21 := sumMatrix(m2, m4)
	c22 := sumMatrix(subtractMatrix(sumMatrix(m1, m3), m2), m6)

	// Combine submatrices into result
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < newSize; i++ {
		for j := 0; j < newSize; j++ {
			result[i][j] = c11[i][j]
			result[i][j+newSize] = c12[i][j]
			result[i+newSize][j] = c21[i][j]
			result[i+newSize][j+newSize] = c22[i][j]
		}
	}

	return result
}
