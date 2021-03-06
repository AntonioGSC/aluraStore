package routes

import (
	"net/http"

	"github.com/AntonioGSC/AluraStore/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create", controllers.Create)
	http.HandleFunc("/insert", controllers.Insert)
}
