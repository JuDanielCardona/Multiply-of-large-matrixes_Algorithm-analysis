import numpy as np
import threading
import time
import math

def Multiply(B, C):
    print("Método: RowByRowEnhanced")

    start_time = time.time_ns()
    size = len(C)
    A = np.zeros((size, size), dtype=np.int64)

    bsize = max(1, int(math.sqrt(size)))  # Calcula bsize como la raíz cuadrada del tamaño de la matriz

    def worker(start, end):
        for i1 in range(start, end, bsize):
            for j1 in range(0, size, bsize):
                for k1 in range(0, size, bsize):
                    for i in range(i1, min(i1 + bsize, size)):
                        for j in range(j1, min(j1 + bsize, size)):
                            for k in range(k1, min(k1 + bsize, size)):
                                A[i][k] += B[i][j] * C[j][k]

    threads = []
    mid = size // 2

    thread1 = threading.Thread(target=worker, args=(0, mid))
    thread2 = threading.Thread(target=worker, args=(mid, size))

    thread1.start()
    thread2.start()

    threads.append(thread1)
    threads.append(thread2)

    for thread in threads:
        thread.join()

    end_time = time.time_ns()
    execution_time = end_time - start_time
    return A, execution_time
