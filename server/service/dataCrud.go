package service

import (
	"errors"
	"fmt"

	"github.com/nj-jay/httpServer/global"

	"github.com/nj-jay/httpServer/models"
)

type Status int //响应值 http code 200 OK

func QueryData() []models.BookData {

	//rows, err := db.Query(database.QueryAllData, (pageSize-1)*8)

	var books []models.BookData

	global.GMD_DB.Find(&books)

	//fmt.Println(books)

	return books
}

//通过name, types, author字段进行模糊匹配
//2020.11.19 00:49更新
func QueryDataByName(name string) []models.BookData {

	fmt.Println(name)

	//rows, err := db.Query(database.QueryDataByName, name)

	var books []models.BookData

	global.GMD_DB.Where("name=?", name).Find(&books)

	return books

}

//根据前端传来的id 进行相应的删除操作　并返回给controllers层响应代码
func DeleteSingleDataById(id string) error {

	//_, err := db.Exec(database.DeleteSingleDataById, id)

	result := global.GMD_DB.Delete(&models.BookData{}, id)

	if result.Error != nil || result.RowsAffected == 0 {

		return errors.New("error")

	}

	return nil

}

func UpdateSingleDataById(id int, price float32) Status {


	//fmt.Println(id, name, types, author, price, addTime, id)

	//_, err := db.Exec(database.UpdateSingleDataById, name, types, author, price, addTime, id)

	result := global.GMD_DB.Model(&models.BookData{}).Where("id=?", id).Update("price", price)

	if result.Error != nil {

		return 401

	}

	return 200

}

func PostSingleData(name, types, author string, price float32, addTime string) Status {

	//_, err := db.Exec(database.PostSingleData, name, types, author, price, addTime)

	book := &models.BookData{

		Name:    name,

		Types:   types,

		Author:  author,

		Price:   price,

		AddTime: addTime,

	}

	result := global.GMD_DB.Create(&book)

	if result != nil {

		return 401

	}

	return 200
}
