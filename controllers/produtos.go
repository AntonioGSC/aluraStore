package controllers

import (
	"net/http"
	"text/template"

	"github.com/AntonioGSC/AluraStore/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}
