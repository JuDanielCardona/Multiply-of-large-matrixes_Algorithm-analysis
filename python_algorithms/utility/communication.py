from flask import request
from utility.algorithmExecutor import init_execution

class Datos:
    def __init__(self, testID, matrixA, matrixB):
        self.TestID = testID
        self.MatrixA = matrixA
        self.MatrixB = matrixB

def Init_PythonExecution(request):
    data = request.get_json()

    if not data:
        return "Error al leer el cuerpo de la solicitud", 400

    try:
        datos = Datos(data['TestID'], data['MatrixA'], data['MatrixB'])
    except KeyError:
        return "Error al decodificar los datos JSON", 400

    test = datos.TestID
    matrizA = datos.MatrixA
    matrizB = datos.MatrixB

    # Llama a la función init_execution con las matrices y el test ID
    init_execution(matrizA, matrizB, test)

    return "OK - Correctly recorded execution times", 200
