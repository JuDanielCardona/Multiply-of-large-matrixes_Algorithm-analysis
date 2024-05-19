import numpy as np
import threading
import time

def Multiply(B, C):
    print("Método: RowByRowParallel")

    start_time = time.time_ns()

    size = len(B)
    A = np.zeros((size, size), dtype=np.int64)

    bsize = size // 2
    threads = []

    for i1 in range(0, size, bsize):
        for j1 in range(0, size, bsize):
            for k1 in range(0, size, bsize):
                thread = threading.Thread(target=worker, args=(i1, j1, k1, A, B, C, bsize, size))
                thread.start()
                threads.append(thread)

    # Esperar a que todos los hilos terminen
    for thread in threads:
        thread.join()

    end_time = time.time_ns()
    execution_time = end_time - start_time

    return A, execution_time

def worker(i1, j1, k1, A, B, C, bsize, size):
    for i in range(i1, min(i1 + bsize, size)):
        for j in range(j1, min(j1 + bsize, size)):
            for k in range(k1, min(k1 + bsize, size)):
                A[i][k] += B[i][j] * C[j][k]
