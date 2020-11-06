package controllers

import (

	"github.com/gin-gonic/gin"

	"github.com/nj-jay/httpServer/service"

)

func QueryAllData(c *gin.Context) {

	allData := service.QueryData()

	c.IndentedJSON(200, allData)

}