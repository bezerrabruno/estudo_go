package controllers

import (
	"api/models"
	"github.com/gin-gonic/gin"
)

func BuscarNovidades(c *gin.Context) {
	c.JSON(200, models.Novidades)
}

func BuscarNovidadePorVersao(c *gin.Context) {
	versao := c.Param("versao")

	for _, item := range models.Novidades {
		if item.Versao == versao {
			c.JSON(200, item)
		}
	}
}
