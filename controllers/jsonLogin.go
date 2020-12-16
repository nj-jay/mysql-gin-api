package controllers

import (

	"github.com/gin-gonic/gin"

	"net/http"

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
// @Header
// @Router /api/v2/trueLogin [post]
func TrueLogin(c *gin.Context) {

	username := c.PostForm("username")

	password := c.PostForm("password")

	tokenString, err := service.TrueLogin(username, password)

	if err != nil {

		c.IndentedJSON(401, gin.H{

			"code": 401,

			"msg": tokenString,

		})

	} else {

		c.IndentedJSON(http.StatusOK, gin.H {

			"code" : 200,

			"msg" : "success",

			"data" : gin.H{

				"token" : tokenString,
			},

		})

	}


}


// @Summary 注册接口
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {} json "{"code":200, "msg":"success add user"}"
// @Failure 401 {} json "{"code":401, "msg":"add user error"}"
// @Router /api/v2/login/add [post]
func PostLogin(c *gin.Context) {

	username := c.PostForm("username")

	password := c.PostForm("password")

	status := service.PostLogin(username, password)

	if status == 200 {

		c.IndentedJSON(http.StatusOK, gin.H{

			"code" : 200,

			"msg" : "success add user",

		})

	} else {

		c.IndentedJSON(http.StatusOK, gin.H {

			"code" : 401,

			"msg" : "add user err",

		})

	}
}