import algorithms.ColumnByColumnParallel 
import algorithms.ColumnByColumnSequential 
import algorithms.NaivLoopUnrollingFour 
import algorithms.NaivLoopUnrollingTwo 
import algorithms.NaivOnArray 
import algorithms.RowByColumnEnhanced 
import algorithms.RowByColumnParallel 
import algorithms.RowByColumnSequential 
import algorithms.RowByRowEnhanced 
import algorithms.RowByRowParallel 
import algorithms.RowByRowSequential 
import algorithms.StrassenNaiv 
import algorithms.StrassenWinograd 
import algorithms.WinogradOriginal 
import algorithms.WinogradScaled
from database.connection import register_info

def init_execution(matrizA, matrizB, test):
    
    result, time = algorithms.NaivOnArray.Multiply(matrizA, matrizB)
    register_info(test, "Python", "NaivOnArray", time)
    print(time, "\n\n")

    result, time = algorithms.NaivLoopUnrollingTwo.Multiply(matrizA, matrizB)
    register_info(test, "Python", "NaivLoopUnrollingTwo", time)
    print(time, "\n\n")

    result, time = algorithms.NaivLoopUnrollingFour.Multiply(matrizA, matrizB)
    register_info(test, "Python", "NaivLoopUnrollingFour", time)
    print(time, "\n\n")

    result, time = algorithms.WinogradOriginal.Multiply(matrizA, matrizB)
    register_info(test, "Python", "WinogradOriginal", time)
    print(time, "\n\n")

    result, time = algorithms.WinogradScaled.Multiply(matrizA, matrizB)
    register_info(test, "Python", "WinogradScaled", time)
    print(time, "\n\n")

    result, time = algorithms.StrassenNaiv.Multiply(matrizA, matrizB)
    register_info(test, "Python", "StrassenNaiv", time)
    print(time, "\n\n")

    result, time = algorithms.StrassenWinograd.Multiply(matrizA, matrizB)
    register_info(test, "Python", "StrassenWinograd", time)
    print(time, "\n\n")

    result, time = algorithms.RowByColumnSequential.Multiply(matrizA, matrizB)
    register_info(test, "Python", "RowByColumnSequential", time)
    print(time, "\n\n")

    result, time = algorithms.RowByColumnParallel.Multiply(matrizA, matrizB)
    register_info(test, "Python", "RowByColumnParallel", time)
    print(time, "\n\n")

    result, time = algorithms.RowByColumnEnhanced.Multiply(matrizA, matrizB)
    register_info(test, "Python", "RowByColumnEnhanced", time)
    print(time, "\n\n")

    result, time = algorithms.RowByRowSequential.Multiply(matrizA, matrizB)
    register_info(test, "Python", "RowByRowSequential", time)
    print(time, "\n\n")

    result, time = algorithms.RowByRowParallel.Multiply(matrizA, matrizB)
    register_info(test, "Python", "RowByRowParallel", time)
    print(time, "\n\n")

    result, time = algorithms.RowByRowEnhanced.Multiply(matrizA, matrizB)
    register_info(test, "Python", "RowByRowEnhanced", time)
    print(time, "\n\n")

    result, time = algorithms.ColumnByColumnSequential.Multiply(matrizA, matrizB)
    register_info(test, "Python", "ColumnByColumnSequential", time)
    print(time, "\n\n")

    result, time = algorithms.ColumnByColumnParallel.Multiply(matrizA, matrizB)
    register_info(test, "Python", "ColumnByColumnParallel", time)
    print(time, "\n\n")

