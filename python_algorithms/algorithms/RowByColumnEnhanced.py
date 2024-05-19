import numpy as np
import threading
import time

def Multiply(B, C):
    print("Método: RowByColumnEnhanced")

    start_time = time.time_ns()
    size = len(C)
    A = np.zeros((size, size), dtype=np.int64)

    def worker(start, end):
        for i in range(start, end):
            for j in range(size):
                for k in range(size):
                    A[i][j] += B[i][k] * C[k][j]

    threads = []
    mid = size // 2

    # Create two threads
    thread1 = threading.Thread(target=worker, args=(0, mid))
    thread2 = threading.Thread(target=worker, args=(mid, size))

    # Start threads
    thread1.start()
    thread2.start()

    # Append threads to the list
    threads.append(thread1)
    threads.append(thread2)

    # Wait for both threads to complete
    for thread in threads:
        thread.join()

    end_time = time.time_ns()
    execution_time = end_time - start_time
    return A, execution_time
