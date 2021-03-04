package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaBD() *sql.DB {
	conexao := "user=postgres dbname=alura_store password=203502 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id, Quantidade  int
	Nome, Descricao string
	Preco           float64
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaBD()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
	defer db.Close()
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaBD()
	selectTodosProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
