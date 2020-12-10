package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nj-jay/httpServer/controllers"
)

func LoadLogin(e *gin.Engine) {

	e.GET("/api/v2/login", controllers.QueryAllLogin)

	e.POST("/api/v2/trueLogin", controllers.TrueLogin)

	e.POST("/api/v2/login/add", controllers.PostLogin)



}