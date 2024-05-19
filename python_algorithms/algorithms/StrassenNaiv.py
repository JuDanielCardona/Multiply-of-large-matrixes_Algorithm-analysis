import numpy as np
import time

def Multiply(matrixA, matrixB):
    print("Método: StrassenNaive")
    start_time = time.time_ns()
    result = multiplication(matrixA, matrixB)
    end_time = time.time_ns()
    execution_time = end_time - start_time
    return result, execution_time

def add_matrix(matrixA, matrixB, matrixC, split_index):
    for i in range(split_index):
        for j in range(split_index):
            matrixC[i][j] = matrixA[i][j] + matrixB[i][j]

def multiplication(matrixA, matrixB):
    col1 = len(matrixA[0])
    row1 = len(matrixA)
    col2 = len(matrixB[0])
    row2 = len(matrixB)

    if col1 != row2:
        return [[0]]

    result_matrix = np.zeros((row1, col2), dtype=np.int64)

    if col1 == 1:
        result_matrix[0][0] = matrixA[0][0] * matrixB[0][0]
    else:
        split_index = col1 // 2

        result_matrix00 = np.zeros((split_index, split_index), dtype=np.int64)
        result_matrix01 = np.zeros((split_index, split_index), dtype=np.int64)
        result_matrix10 = np.zeros((split_index, split_index), dtype=np.int64)
        result_matrix11 = np.zeros((split_index, split_index), dtype=np.int64)

        a00 = np.zeros((split_index, split_index), dtype=np.int64)
        a01 = np.zeros((split_index, split_index), dtype=np.int64)
        a10 = np.zeros((split_index, split_index), dtype=np.int64)
        a11 = np.zeros((split_index, split_index), dtype=np.int64)
        b00 = np.zeros((split_index, split_index), dtype=np.int64)
        b01 = np.zeros((split_index, split_index), dtype=np.int64)
        b10 = np.zeros((split_index, split_index), dtype=np.int64)
        b11 = np.zeros((split_index, split_index), dtype=np.int64)

        for i in range(split_index):
            for j in range(split_index):
                a00[i][j] = matrixA[i][j]
                a01[i][j] = matrixA[i][j+split_index]
                a10[i][j] = matrixA[i+split_index][j]
                a11[i][j] = matrixA[i+split_index][j+split_index]
                b00[i][j] = matrixB[i][j]
                b01[i][j] = matrixB[i][j+split_index]
                b10[i][j] = matrixB[i+split_index][j]
                b11[i][j] = matrixB[i+split_index][j+split_index]

        add_matrix(multiplication(a00, b00), multiplication(a01, b10), result_matrix00, split_index)
        add_matrix(multiplication(a00, b01), multiplication(a01, b11), result_matrix01, split_index)
        add_matrix(multiplication(a10, b00), multiplication(a11, b10), result_matrix10, split_index)
        add_matrix(multiplication(a10, b01), multiplication(a11, b11), result_matrix11, split_index)

        for i in range(split_index):
            for j in range(split_index):
                result_matrix[i][j] = result_matrix00[i][j]
                result_matrix[i][j+split_index] = result_matrix01[i][j]
                result_matrix[i+split_index][j] = result_matrix10[i][j]
                result_matrix[i+split_index][j+split_index] = result_matrix11[i][j]

    return result_matrix
