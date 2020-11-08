package controllers

import (

	"github.com/gin-gonic/gin"

	"github.com/nj-jay/httpServer/service"
)

func QueryAllLogin(c *gin.Context) {

	allLogin := service.QueryLogin()

	c.IndentedJSON(200, allLogin)

}