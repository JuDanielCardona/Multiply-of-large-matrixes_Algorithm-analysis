import numpy as np
import time

def Multiply(A, B):
    print("Método: ColumnByColumnSequential")

    start_time = time.time_ns()

    rowsA, colsA, colsB = len(A), len(A[0]), len(B[0])

    result = np.zeros((rowsA, colsB), dtype=np.int64)

    for j in range(colsB):
        for i in range(rowsA):
            for k in range(colsA):
                result[i][j] += A[i][k] * B[k][j]

    end_time = time.time_ns()
    execution_time = end_time - start_time
    return result, execution_time
