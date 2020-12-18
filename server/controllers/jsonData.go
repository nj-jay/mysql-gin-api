package controllers

import (

	"github.com/nj-jay/httpServer/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nj-jay/httpServer/service"

	"github.com/spf13/cast"

)

func QueryAllData(c *gin.Context) {

	allData := service.QueryData()

	c.IndentedJSON(200, models.NewGetBooksSuccess(allData))

}

func QueryDataByName(c *gin.Context) {

	name := c.PostForm("name")

	dataByName := service.QueryDataByName(name)

	c.IndentedJSON(200, dataByName)

}

func DeleteSingleDataById(c *gin.Context) {

	id := c.PostForm("id")

	err := service.DeleteSingleDataById(id)

	if err != nil {

		c.IndentedJSON(http.StatusOK, models.NewDeleteBookByIdErr(id))

	} else {

		c.IndentedJSON(http.StatusOK, models.NewDeleteBookByIdSuccess(id))

	}
}

func UpdateSingleDataById(c *gin.Context) {

	id := c.PostForm("id")

	price := c.PostForm("price")

	idc := cast.ToInt(id)

	pricec := cast.ToFloat32(price)

	status := service.UpdateSingleDataById(idc, pricec)

	if status == 200 {

		c.String(200, "success")

	} else {

		c.String(404, "error")

	}
}

func PostSingleData(c *gin.Context) {

	name := c.PostForm("name")

	types := c.PostForm("types")

	author := c.PostForm("author")

	price := c.PostForm("price")

	addTime := c.PostForm("addTime")

	pricec := cast.ToFloat32(price)

	status := service.PostSingleData(name, types, author, pricec, addTime)

	if status == 200 {

		c.String(200, "success")

	} else {

		c.String(404, "error")

	}

}