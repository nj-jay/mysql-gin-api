package routers

import (

	"github.com/gin-gonic/gin"

	"github.com/nj-jay/httpServer/middlewares"

	"github.com/nj-jay/httpServer/controllers"
)

func Router() *gin.Engine {

	r := gin.New()

	r.Use(middlewares.Cors())

	r.GET("/api/v1/books", controllers.QueryAllData)

	return r

}