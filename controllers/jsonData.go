package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nj-jay/httpServer/service"
)

func QueryAllData(c *gin.Context) {

	page := c.Param("page")

	allData := service.QueryData(page)

	fmt.Println(allData)

	c.IndentedJSON(200, allData)

}

func QueryDataByName(c *gin.Context) {

	name := c.PostForm("name")

	dataByName := service.QueryDataByName(name)

	fmt.Println(dataByName)

	c.IndentedJSON(200, dataByName)

}

func DeleteSingleDataById(c *gin.Context) {

	id := c.PostForm("id")

	status := service.DeleteSingleDataById(id)

	if status == 200 {

		c.String(200, "success")

	} else if status == 404 {

		c.String(404, "error")

	}
}

func UpdateSingleDataById(c *gin.Context) {

	id := c.PostForm("id")

	name := c.PostForm("name")

	types := c.PostForm("types")

	author := c.PostForm("author")

	price := c.PostForm("price")

	addTime := c.PostForm("addTime")

	status := service.UpdateSingleDataById(name, types, author, price, addTime, id)

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

	status := service.PostSingleData(name, types, author, price, addTime)

	if status == 200 {

		c.String(200, "success")

	} else {

		c.String(404, "error")

	}

}