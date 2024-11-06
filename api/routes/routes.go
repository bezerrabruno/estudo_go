package routes

import (
	"api/controllers"
	"github.com/gin-gonic/gin"
)

func StartRotas(r *gin.Engine) {
	r.GET("/", controllers.Api)
	r.GET("/api/v1/novidades", controllers.BuscarNovidades)
	r.GET("/api/v1/novidades/:versao", controllers.BuscarNovidadePorVersao)
}
