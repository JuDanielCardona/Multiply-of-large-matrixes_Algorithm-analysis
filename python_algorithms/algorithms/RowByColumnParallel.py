import numpy as np
import threading
import time

def Multiply(A, B):
    print("Método: RowByColumnParallel")

    start_time = time.time_ns()

    size = len(A)
    C = np.zeros((size, size), dtype=np.int64)

    def calculate_element(row, col):
        sum_val = 0
        for k in range(size):
            sum_val += A[row][k] * B[k][col]
        C[row][col] = sum_val

    threads = []
    for i in range(size):
        for j in range(size):
            thread = threading.Thread(target=calculate_element, args=(i, j))
            thread.start()
            threads.append(thread)

    for thread in threads:
        thread.join()

    end_time = time.time_ns()
    execution_time = end_time - start_time
    return C, execution_time
