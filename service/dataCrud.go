package service

import (

	"fmt"

	"github.com/spf13/cast"

	"github.com/nj-jay/httpServer/database"

	"github.com/nj-jay/httpServer/models"

)

type BookData models.BookData

type Status int //响应值 http code 200 OK

//与数据库进行交互,根据前端传入的page值,
//返回相应的数据传给controllers层(实现分页的效果)
func QueryData(page string) []BookData {

	db, _ := database.Connect()

	defer db.Close()

	pageSize := cast.ToInt(page)

	fmt.Println(pageSize)

	rows, err := db.Query(database.QueryAllData, (pageSize-1)*8)

	sliceBook := make([]BookData, 0)

	bookData := new(BookData)

	if err != nil {

		fmt.Println(err)
	}

	for rows.Next() {

		err := rows.Scan(&bookData.Id, &bookData.Name, &bookData.Types,
			&bookData.Author, &bookData.Price, &bookData.AddTime)

		if err != nil {

			fmt.Println("error")

		}

		sliceBook = append(sliceBook, *bookData)

	}

	return sliceBook

}

//通过name, types, author字段进行模糊匹配
//2020.11.19 00:49更新
func QueryDataByName(name string) []BookData {

	db, _ := database.Connect()

	defer db.Close()

	fmt.Println(name)

	rows, err := db.Query(database.QueryDataByName, name)

	sliceBook := make([]BookData, 0)

	bookData := new(BookData)

	if err != nil {

		fmt.Println(err)
	}

	for rows.Next() {

		err := rows.Scan(&bookData.Id, &bookData.Name, &bookData.Types,
			&bookData.Author, &bookData.Price, &bookData.AddTime)

		if err != nil {

			fmt.Println("error")

		}

		sliceBook = append(sliceBook, *bookData)
	}

	return sliceBook

}

//根据前端传来的id 进行相应的删除操作　并返回给controllers层响应代码
func DeleteSingleDataById(id string) Status {

	db, _ := database.Connect()

	defer db.Close()

	fmt.Println(id)

	_, err := db.Exec(database.DeleteSingleDataById, id)

	if err != nil {

		return 404

	} else {

		return 200
	}
}

func UpdateSingleDataById(name, types, author, price , addTime, id string) Status {

	db, _ := database.Connect()

	defer db.Close()

	fmt.Println(id, name, types, author, price, addTime, id)

	_, err := db.Exec(database.UpdateSingleDataById, name, types, author, price, addTime, id)

	if err != nil {

		return 404

	} else {

		return 200
	}
}

func PostSingleData(name, types, author, price, addTime string) Status {

	db, _ := database.Connect()

	defer db.Close()

	_, err := db.Exec(database.PostSingleData, name, types, author, price, addTime)

	if err != nil {

		return  404

	} else {

		return 200
	}
}
