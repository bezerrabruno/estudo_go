package controllers

import "github.com/gin-gonic/gin"

func Api(c *gin.Context) {
	c.JSON(200, gin.H{"mensagem": "pong"})
}
