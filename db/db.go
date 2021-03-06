package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func ConectaBD() *sql.DB {
	conexao := os.ExpandEnv("user=postgres dbname=alura_store password=${DB_PASSWORD} host=localhost sslmode=disable")
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
