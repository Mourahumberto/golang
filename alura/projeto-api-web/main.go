package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=neto dbname=netodb password=neto123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Quant     int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}
	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quant = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}