package algorithms

import (
	"fmt"
	"sync"
	"time"
)

func RowByRowParallel(B, C [][]int) ([][]int, int64) {
	fmt.Println("Método: RowByRowParallel")

	startTime := time.Now().UnixNano()

	size := len(B)
	A := make([][]int, size)
	for i := range A {
		A[i] = make([]int, size)
	}

	bsize := size / 2
	var wg sync.WaitGroup

	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				wg.Add(1)
				go func(i1, j1, k1 int) {
					defer wg.Done()
					for i := i1; i < i1+bsize && i < size; i++ {
						for j := j1; j < j1+bsize && j < size; j++ {
							for k := k1; k < k1+bsize && k < size; k++ {
								A[i][k] += B[i][j] * C[j][k]
							}
						}
					}
				}(i1, j1, k1)
			}
		}
	}

	wg.Wait()

	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime

	return A, executionTime
}
