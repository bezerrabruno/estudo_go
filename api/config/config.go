package config

import (
	"api/models"
	"api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartConfig() *gin.Engine {
	fmt.Println("Criando modelos")
	models.Novidades = []models.NovidadeModel{
		{Titulo: "Lançamento", Notas: "Finalmente é lançado novo chat!!!", Versao: "0.1"},
		{Titulo: "Versao estavel", Notas: "Nova versao estavel disponivel para celulares.", Versao: "0.2"},
	}

	fmt.Println("Configurando o servidor")
	r := gin.Default()

	fmt.Println("Configurando rotas")
	routes.StartRotas(r)

	return r
}
