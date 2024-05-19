import numpy as np
import time

def Multiply(A, B):
    print("Método: naiveLoopUnrollingFour")

    start_time = time.time_ns()
    N = len(A)
    M = len(B[0])
    P = len(B)
    result = np.zeros((N, M), dtype=np.int64)

    for i in range(N):
        for j in range(M):
            aux = 0
            if P % 4 == 0:
                for k in range(0, P, 4):
                    aux += (A[i][k] * B[k][j] +
                            A[i][k+1] * B[k+1][j] +
                            A[i][k+2] * B[k+2][j] +
                            A[i][k+3] * B[k+3][j])
            elif P % 4 == 1:
                PP = P - 1
                for k in range(0, PP, 4):
                    aux += (A[i][k] * B[k][j] +
                            A[i][k+1] * B[k+1][j] +
                            A[i][k+2] * B[k+2][j] +
                            A[i][k+3] * B[k+3][j])
                aux += A[i][PP] * B[PP][j]
            elif P % 4 == 2:
                PP = P - 2
                PPP = P - 1
                for k in range(0, PP, 4):
                    aux += (A[i][k] * B[k][j] +
                            A[i][k+1] * B[k+1][j] +
                            A[i][k+2] * B[k+2][j] +
                            A[i][k+3] * B[k+3][j])
                aux += (A[i][PP] * B[PP][j] +
                        A[i][PPP] * B[PPP][j])
            else:
                PP = P - 3
                PPP = P - 2
                PPPP = P - 1
                for k in range(0, PP, 4):
                    aux += (A[i][k] * B[k][j] +
                            A[i][k+1] * B[k+1][j] +
                            A[i][k+2] * B[k+2][j] +
                            A[i][k+3] * B[k+3][j])
                aux += (A[i][PP] * B[PP][j] +
                        A[i][PPP] * B[PPP][j] +
                        A[i][PPPP] * B[PPPP][j])
            result[i][j] = aux

    end_time = time.time_ns()
    execution_time = end_time - start_time
    return result, execution_time
