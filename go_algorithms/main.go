package main

import (
	"fmt"
	"go_algorithms/database"
	"go_algorithms/utility"
	"log"
	"net/http"
)

func main() {
	database.Connection()
	http.HandleFunc("/", utility.Init_GolangExecution)
	fmt.Println("Server is listening on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
