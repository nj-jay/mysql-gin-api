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

	r.GET("/api/v1/book/:id", controllers.QuerySingleDataById)

	r.GET("/api/v1/login", controllers.QueryAllLogin)

	r.DELETE("/api/v1/book/delete", controllers.DeleteSingleDataById)

	r.PUT("/api/v1/book/update/:id", controllers.UpdateSingleDataById)

	r.POST("/api/v1/book/add", controllers.PostSingleData)

	r.POST("/api/v1/login", controllers.TrueLogin)

	return r

}