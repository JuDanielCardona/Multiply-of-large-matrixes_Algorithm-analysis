import numpy as np
import threading
import time


def Multiply(A, B):
        print("Método: ColumnByColumnParallel")

        start_time = time.time_ns()

        rowsA, colsA, colsB = len(A), len(A[0]), len(B[0])

        result = np.zeros((rowsA, colsB), dtype=np.int64)

        def calculate_element(row, col):
            sum_val = 0
            for k in range(colsA):
                sum_val += A[row][k] * B[k][col]
            result[row][col] = sum_val

        threads = []
        for j in range(colsB):
            for i in range(rowsA):
                thread = threading.Thread(target=calculate_element, args=(i, j))
                thread.start()
                threads.append(thread)

        for thread in threads:
            thread.join()

        end_time = time.time_ns()
        execution_time = end_time - start_time
        return result, execution_time
