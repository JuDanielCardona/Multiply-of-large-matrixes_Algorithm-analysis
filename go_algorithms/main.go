package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go_algorithms/utility"
)

func main() {

	fileRoute := "sources/matrizA.txt"

	matrixA, err := readMatrix(fileRoute)
	if err != nil {
		fmt.Printf("Error al leer el archivo %s: %v\n", fileRoute, err)
		return
	}

	fileRoute = "sources/matrizB.txt"

	matrixB, err := readMatrix(fileRoute)
	if err != nil {
		fmt.Printf("Error al leer el archivo %s: %v\n", fileRoute, err)
		return
	}

	utility.Init_GoExecution(matrixA, matrixB)

}

func readMatrix(filename string) ([][]int, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		row := make([]int, len(fields))

		for i, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			row[i] = num
		}

		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}
