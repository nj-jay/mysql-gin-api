package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nj-jay/httpServer/global"
	"github.com/nj-jay/httpServer/initialization"
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

	global.GMD_DB = initalization.GormMysql()

	initalization.MysqlTables(global.GMD_DB)

	r := gin.Default()

	r.Use(middlewares.Cors())

	routers.LoadLogin(r)

	routers.LoadBooks(r)

	err := http.ListenAndServeTLS(":8082", "conf/api.nj-jay.com_bundle.crt", "conf/api.nj-jay.com.key", r)

	//err := r.Run(":8080")

	if err != nil {

		log.Fatal("run err")

	}
}
