package utility

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Datos struct {
	TestID  string  `json:"TestID"`
	MatrixA [][]int `json:"MatrixA"`
	MatrixB [][]int `json:"MatrixB"`
}

func Init_GolangExecution(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	var datos Datos
	if err := json.Unmarshal(body, &datos); err != nil {
		http.Error(w, "Error al decodificar los datos JSON", http.StatusBadRequest)
		return
	}

	test := datos.TestID
	matrizA := datos.MatrixA
	matrizB := datos.MatrixB

	Init_execution(matrizA, matrizB, test)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK - Correctly recorded execution times"))
}
