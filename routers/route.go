package routers

import (

	"github.com/gin-gonic/gin"

	"github.com/nj-jay/httpServer/middlewares"

	"github.com/nj-jay/httpServer/controllers"
)

func Router() *gin.Engine {

	r := gin.New()

	r.Use(middlewares.Cors())

	r.GET("/api/v2/books/:page", controllers.QueryAllData)

	r.POST("/api/v2/book/search", controllers.QueryDataByName)

	r.GET("/api/v2/login", controllers.QueryAllLogin)

	r.POST("/api/v2/book/delete", controllers.DeleteSingleDataById)

	r.POST("/api/v2/book/update", controllers.UpdateSingleDataById)

	r.POST("/api/v2/book/add", controllers.PostSingleData)

	r.POST("/api/v2/trueLogin", controllers.TrueLogin)

	r.POST("/api/v2/login/add", controllers.PostLogin)

	return r

}
