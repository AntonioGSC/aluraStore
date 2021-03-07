package model

import (
	"github.com/AntonioGSC/AluraStore/db"
)

type Produto struct {
	Id, Quantidade  int
	Nome, Descricao string
	Preco           float64
}

func BuscaProduto(ID string) Produto {
	db := db.ConectaBD()
	produtoBanco, err := db.Query("select * from produtos where id=$1", ID)

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	defer db.Close()
	return produto
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaBD()
	selectTodosProdutos, err := db.Query("select * from produtos order by id asc")

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

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaBD()
	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(ID string) {
	db := db.ConectaBD()
	deletaDoBanco, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletaDoBanco.Exec(ID)
	defer db.Close()
}

func EditaProduto(ID string) Produto {
	produto := BuscaProduto(ID)
	return produto
}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaBD()
	atualizaProdutoBanco, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	atualizaProdutoBanco.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
