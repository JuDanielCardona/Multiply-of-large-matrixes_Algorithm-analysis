import numpy as np
import time

def Multiply(B, C):
    print("Método: RowByRowSequential")

    start_time = time.time_ns()

    size = len(C)
    A = np.zeros((size, size), dtype=np.int64)

    for i in range(size):
        for j in range(size):
            for k in range(size):
                A[i][k] += B[i][j] * C[j][k]

    end_time = time.time_ns()
    execution_time = end_time - start_time

    return A, execution_time
