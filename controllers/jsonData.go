package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nj-jay/httpServer/models"
	"github.com/nj-jay/httpServer/service"
)

func QueryAllData(c *gin.Context) {

	allData := service.QueryData()

	c.IndentedJSON(200, allData)

}

func QuerySingleDataById(c *gin.Context) {

	id := c.Param("id")

	singleData, status := service.QuerySingleDataById(id)

	if status == 404 {

		c.IndentedJSON(404, models.NewError())


	} else {

		c.IndentedJSON(200, singleData)
	}
}

func DeleteSingleDataById(c *gin.Context) {

	id := c.Query("id")

	status := service.DeleteSingleDataById(id)

	if status == 200 {

		c.IndentedJSON(200, models.NewSucceedDelete())

	} else if status == 404 {


		c.IndentedJSON(404, models.NewErrorDelete())

	}
}

func UpdateSingleDataById(c *gin.Context) {

	id := c.Param("id")

	price := c.PostForm("price")

	status := service.UpdateSingleDataById(id, price)

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