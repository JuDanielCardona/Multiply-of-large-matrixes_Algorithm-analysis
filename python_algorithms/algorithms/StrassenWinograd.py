import numpy as np
import time

def Multiply(matrixA, matrixB):
    print("Método: strassenWinograd")
    
    start_time = time.time_ns()
    result = multiplication_sw(matrixA, matrixB)
    end_time = time.time_ns()
    execution_time = end_time - start_time
    return result, execution_time

def sum_matrix(a, b):
    n = len(a)
    result = np.zeros((n, n), dtype=np.int64)
    for i in range(n):
        for j in range(n):
            result[i][j] = a[i][j] + b[i][j]
    return result

def subtract_matrix(a, b):
    n = len(a)
    result = np.zeros((n, n), dtype=np.int64)
    for i in range(n):
        for j in range(n):
            result[i][j] = a[i][j] - b[i][j]
    return result

def multiplication_sw(a, b):
    n = len(a)

    if n == 1:
        return [[a[0][0] * b[0][0]]]

    new_size = n // 2

    a11 = np.array([row[:new_size] for row in a[:new_size]], dtype=np.int64)
    a12 = np.array([row[new_size:] for row in a[:new_size]], dtype=np.int64)
    a21 = np.array([row[:new_size] for row in a[new_size:]], dtype=np.int64)
    a22 = np.array([row[new_size:] for row in a[new_size:]], dtype=np.int64)
    b11 = np.array([row[:new_size] for row in b[:new_size]], dtype=np.int64)
    b12 = np.array([row[new_size:] for row in b[:new_size]], dtype=np.int64)
    b21 = np.array([row[:new_size] for row in b[new_size:]], dtype=np.int64)
    b22 = np.array([row[new_size:] for row in b[new_size:]], dtype=np.int64)

    m1 = multiplication_sw(sum_matrix(a11, a22), sum_matrix(b11, b22))
    m2 = multiplication_sw(sum_matrix(a21, a22), b11)
    m3 = multiplication_sw(a11, subtract_matrix(b12, b22))
    m4 = multiplication_sw(a22, subtract_matrix(b21, b11))
    m5 = multiplication_sw(sum_matrix(a11, a12), b22)
    m6 = multiplication_sw(subtract_matrix(a21, a11), sum_matrix(b11, b12))
    m7 = multiplication_sw(subtract_matrix(a12, a22), sum_matrix(b21, b22))

    c11 = sum_matrix(subtract_matrix(sum_matrix(m1, m4), m5), m7)
    c12 = sum_matrix(m3, m5)
    c21 = sum_matrix(m2, m4)
    c22 = sum_matrix(subtract_matrix(sum_matrix(m1, m3), m2), m6)

    result = np.zeros((n, n), dtype=np.int64)
    for i in range(new_size):
        for j in range(new_size):
            result[i][j] = c11[i][j]
            result[i][j + new_size] = c12[i][j]
            result[i + new_size][j] = c21[i][j]
            result[i + new_size][j + new_size] = c22[i][j]

    return result
