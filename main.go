package main

import (
	"net/http"
	"os"

	"github.com/AntonioGSC/AluraStore/routes"
)

func main() {
	// Set Environment Variables
	os.Setenv("DB_PASSWORD", "*SUA SENHA AQUI*")

	//Starting server
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
