package routes

import (
	"net/http"

	"github.com/AntonioGSC/AluraStore/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
