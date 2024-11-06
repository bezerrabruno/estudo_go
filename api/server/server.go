package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer(r *gin.Engine) {
	fmt.Println("Iniciando o servidor...")

	err := r.Run(":9010")

	if err != nil {
		panic(err)
	}
}
