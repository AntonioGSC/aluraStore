package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Maneira", Preco: 49.99, Quantidade: 19},
		{"Calça", "Azul", 99.99, 11},
		{"Notebook", "Rápido", 2999.99, 5},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
