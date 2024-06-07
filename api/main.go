package main

import (
	"fmt"

	"github.com/guilhermefparra/api-go-rest/database"
	"github.com/guilhermefparra/api-go-rest/models"
	"github.com/guilhermefparra/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia da primeira"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia da segunda"},
	}
	database.ConectarAoBancoDeDados()
	fmt.Println("Iniciando o servidor REST")
	routes.HandleRequest()
}
