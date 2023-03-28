package main

import (
	"projeto-gin/database"
	"projeto-gin/routes"
)

func main() {
	database.ConectaBancoDeDados()
	routes.HandleRequest()
}
