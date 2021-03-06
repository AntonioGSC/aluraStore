package main

import (
	"net/http"
	"os"

	"github.com/AntonioGSC/AluraStore/routes"
)

func main() {
	// Set Environment Variables
	os.Setenv("DB_PASSWORD", "_SUA SENHA AQUI_")

	//Starting server
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
