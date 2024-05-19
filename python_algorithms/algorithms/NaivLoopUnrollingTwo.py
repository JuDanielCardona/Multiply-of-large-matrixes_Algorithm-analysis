import numpy as np
import time

def Multiply(A, B):
    print("Método: naiveLoopUnrollingTwo")

    start_time = time.time_ns()
    N = len(A)
    M = len(B[0])
    P = len(B)
    result = np.zeros((N, M), dtype=np.int64)

    for i in range(N):
        for j in range(M):
            aux = 0
            if P % 2 == 0:
                for k in range(0, P, 2):
                    aux += A[i][k] * B[k][j] + A[i][k+1] * B[k+1][j]
            else:
                PP = P - 1
                for k in range(0, PP, 2):
                    aux += A[i][k] * B[k][j] + A[i][k+1] * B[k+1][j]
                aux += A[i][PP] * B[PP][j]
            result[i][j] = aux

    end_time = time.time_ns()
    execution_time = end_time - start_time
    return result, execution_time
