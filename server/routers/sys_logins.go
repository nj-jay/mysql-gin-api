package routers

import (

	"github.com/gin-gonic/gin"

	"github.com/nj-jay/httpServer/controllers"

	"github.com/nj-jay/httpServer/middlewares"

	gs "github.com/swaggo/gin-swagger"

	"github.com/swaggo/gin-swagger/swaggerFiles"

)

func LoadLogin(e *gin.Engine) {

	e.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	e.GET("/api/v3/login", middlewares.JWTAuthMiddleware(), controllers.QueryAllLogin)

	e.POST("/api/v3/trueLogin", controllers.TrueLogin)

	e.POST("/api/v3/login/add", controllers.PostLogin)

}