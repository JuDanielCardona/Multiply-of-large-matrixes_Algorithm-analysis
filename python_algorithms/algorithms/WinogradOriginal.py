import numpy as np
import time

def Multiply(A, B):
    print("Método: WinogradOriginal")

    start_time = time.time_ns()

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
            aux += A[i][j] * A[i][j + 1]
        y[i] = aux

    for i in range(N):
        aux = 0
        for j in range(0, gamma, 2):
            aux += B[j][i] * B[j + 1][i]
        z[i] = aux

    if upsilon == 1:
        PP = P - 1
        for i in range(M):
            for k in range(N):
                aux = 0
                for j in range(0, gamma, 2):
                    aux += (A[i][j] + B[j + 1][k]) * (A[i][j + 1] + B[j][k])
                Result[i][k] = aux - y[i] - z[k] + A[i][PP] * B[PP][k]
    else:
        for i in range(M):
            for k in range(N):
                aux = 0
                for j in range(0, gamma, 2):
                    aux += (A[i][j] + B[j + 1][k]) * (A[i][j + 1] + B[j][k])
                Result[i][k] = aux - y[i] - z[k]

    end_time = time.time_ns()
    execution_time = end_time - start_time

    return Result, execution_time
