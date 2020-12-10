package main

import (

	"github.com/gin-gonic/gin"
	"github.com/nj-jay/httpServer/middlewares"
	"github.com/nj-jay/httpServer/routers"
	"log"

)

// @title mysql-gin-api接口文档
// @version 2.0
// @description 接口文档
// @contact.name nj-jay
// @host http://localhost:8080/api/swagger/index.html
func main() {

	//r := routers.Router()

	r := gin.Default()

	r.Use(middlewares.Cors())

	routers.LoadLogin(r)

	routers.LoadBooks(r)

	err := r.Run(":8080")

	if err != nil {

		log.Println("run error")

	}
}