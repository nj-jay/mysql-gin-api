package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nj-jay/httpServer/controllers"
)

func LoadBooks(e *gin.Engine) {


	e.GET("/api/v2/books/:page", controllers.QueryAllData)

	e.POST("/api/v2/book/search", controllers.QueryDataByName)

	e.POST("/api/v2/book/delete", controllers.DeleteSingleDataById)

	e.POST("/api/v2/book/update", controllers.UpdateSingleDataById)

	e.POST("/api/v2/book/add", controllers.PostSingleData)

}