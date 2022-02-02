package main

import (
	"fmt"

	"github.com/JulioBCP/go-rest-api/database"
	"github.com/JulioBCP/go-rest-api/models"
	"github.com/JulioBCP/go-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{ID: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
