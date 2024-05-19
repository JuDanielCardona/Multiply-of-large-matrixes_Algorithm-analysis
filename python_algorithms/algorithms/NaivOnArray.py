import numpy as np
import time

def Multiply(A, B):
    print("Método: naiveOnArray")

    start_time = time.time_ns()
    N = len(A)
    M = len(B[0])
    P = len(B)

    result = np.zeros((N, M), dtype=np.int64)

    for i in range(N):
        for j in range(M):
            aux = 0
            for k in range(P):
                aux += A[i][k] * B[k][j]
            result[i][j] = aux

    end_time = time.time_ns()
    execution_time = end_time - start_time
    return result, execution_time
