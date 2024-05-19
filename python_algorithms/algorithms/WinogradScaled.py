import numpy as np
import math
import time

def Multiply(A, B):
    print("Método: WinogradScaled")

    start_time = time.time_ns()

    N = len(B[0])
    P = len(B)
    M = len(A)

    CopyA = [row.copy() for row in A]
    CopyB = [row.copy() for row in B]

    a = norm_inf(A, N, P)
    b = norm_inf(B, P, M)
    lambda_val = math.floor(0.5 + math.log(b / a) / math.log(4))

    multiply_with_scalar(A, CopyA, N, P, 2 ** lambda_val)
    multiply_with_scalar(B, CopyB, P, M, 2 ** -lambda_val)

    result = winograd(CopyA, CopyB)

    end_time = time.time_ns()
    execution_time = end_time - start_time

    return result, execution_time

def norm_inf(matrix, rows, cols):
    max_norm = 0
    for i in range(rows):
        row_sum = sum(abs(val) for val in matrix[i])
        max_norm = max(max_norm, row_sum)
    return max_norm

def multiply_with_scalar(A, B, rows, cols, scalar):
    for i in range(rows):
        for j in range(cols):
            B[i][j] = int(A[i][j] * scalar)

def winograd(A, B):
    N = len(B[0])
    P = len(B)
    M = len(A)

    Result = np.zeros((M, N), dtype=np.int64)

    upsilon = P % 2
    gamma = P - upsilon
    y = np.zeros(M, dtype=np.int64)
    z = np.zeros(N, dtype=np.int64)

    for i in range(M):
        aux = 0
        for j in range(0, gamma, 2):
            aux += A[i][j] * A[i][j+1]
        y[i] = aux

    for i in range(N):
        aux = 0
        for j in range(0, gamma, 2):
            aux += B[j][i] * B[j+1][i]
        z[i] = aux

    if upsilon == 1:
        PP = P - 1
        for i in range(M):
            for k in range(N):
                aux = 0
                for j in range(0, gamma, 2):
                    aux += (A[i][j] + B[j+1][k]) * (A[i][j+1] + B[j][k])
                Result[i][k] = aux - y[i] - z[k] + A[i][PP] * B[PP][k]
    else:
        for i in range(M):
            for k in range(N):
                aux = 0
                for j in range(0, gamma, 2):
                    aux += (A[i][j] + B[j+1][k]) * (A[i][j+1] + B[j][k])
                Result[i][k] = aux - y[i] - z[k]

    return Result
