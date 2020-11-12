package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nj-jay/httpServer/models"
	"github.com/nj-jay/httpServer/service"
)

func QueryAllData(c *gin.Context) {

	allData := service.QueryData()

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

		c.IndentedJSON(200, "success")

	} else if status == 404 {

		c.IndentedJSON(404, models.NewErrorDelete())

	}
}

func UpdateSingleDataById(c *gin.Context) {

	id := c.PostForm("id")

	name := c.PostForm("name")

	price := c.PostForm("price")

	status := service.UpdateSingleDataById(id, name, price)

	if status == 200 {

		c.IndentedJSON(200, "success")

	} else {

		c.IndentedJSON(404, "error")

	}
}

func PostSingleData(c *gin.Context) {

	name := c.PostForm("name")

	price := c.PostForm("price")

	status := service.PostSingleData(name, price)

	if status == 200 {

		c.IndentedJSON(200, "success")

	} else {

		c.IndentedJSON(404, "error")

	}

}