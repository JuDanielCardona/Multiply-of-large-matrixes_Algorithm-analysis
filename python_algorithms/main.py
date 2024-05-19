import json
from http.server import HTTPServer, BaseHTTPRequestHandler
from flask import request
from utility.algorithmExecutor import init_execution

class Datos:
    def __init__(self, testID, matrixA, matrixB):
        self.TestID = testID
        self.MatrixA = matrixA
        self.MatrixB = matrixB

class RequestHandler(BaseHTTPRequestHandler):
    def do_POST(self):
        content_length = int(self.headers['Content-Length'])
        post_data = json.loads(self.rfile.read(content_length).decode("utf-8"))
        response_message, status_code = Init_PythonExecution(post_data)
        self.send_response(status_code)
        self.send_header('Content-type', 'text/plain')
        self.end_headers()
        self.wfile.write(response_message.encode('utf-8'))

def Init_PythonExecution(data):
    if not data:
        return "Error al leer el cuerpo de la solicitud", 400

    try:
        datos = Datos(data['TestID'], data['MatrixA'], data['MatrixB'])
    except KeyError:
        return "Error al decodificar los datos JSON", 400

    test = datos.TestID
    matrizA = datos.MatrixA
    matrizB = datos.MatrixB

    init_execution(matrizA, matrizB, test)

    return "OK - Correctly recorded execution times", 200

def run(server_class, handler_class, port):
    print(f'init python service, on port: {port}\n\n')
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    httpd.serve_forever() 
    
if __name__ == '__main__':
    run(HTTPServer, RequestHandler, port=8081)
