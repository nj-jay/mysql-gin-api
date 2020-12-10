package main

import (

	"github.com/gin-gonic/gin"
	"github.com/nj-jay/httpServer/middlewares"
	"github.com/nj-jay/httpServer/routers"
	"log"

)

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