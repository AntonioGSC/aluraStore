package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/AntonioGSC/AluraStore/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func Create(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertido, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		model.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertido)
	}
	http.Redirect(w, r, "/", 301)
}
