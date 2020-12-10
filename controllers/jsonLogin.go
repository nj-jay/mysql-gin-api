package controllers

import (

	"github.com/gin-gonic/gin"

	"github.com/nj-jay/httpServer/service"

)

func QueryAllLogin(c *gin.Context) {

	allLogin := service.QueryLogin()

	c.IndentedJSON(200, allLogin)

}

// @Summary 登录接口
// @Produce  json
// @Param username query string true "登录名"
// @Param password query string true "密码"
// @Success 200 {string} json "success"
// @Router /api/v2/trueLogin [post]
func TrueLogin(c *gin.Context) {

	username := c.PostForm("username")

	password := c.PostForm("password")

	status := service.TrueLogin(username, password)

	if status == 200 {

		c.IndentedJSON(200, "success")

	} else {

		c.IndentedJSON(200, "error")

	}
}

func PostLogin(c *gin.Context) {

	username := c.PostForm("username")

	password := c.PostForm("password")

	status := service.PostLogin(username, password)

	if status == 200 {

		c.IndentedJSON(200, "success add user")

	} else {

		c.IndentedJSON(200, "err add user")

	}
}