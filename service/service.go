package service

import (

	"fmt"

	"github.com/nj-jay/httpServer/database"

	"github.com/nj-jay/httpServer/models"

)

type BookData models.BookData

func QueryData() []BookData {

	db, _ := database.Connect()

	defer db.Close()

	rows, err := db.Query(database.QueryAllData)

	sliceBook := make([]BookData, 0)

	bookData := new(BookData)

	if err != nil {

		fmt.Println(err)
	}

	var id int

	var name string

	var price float32

	for rows.Next() {

		err := rows.Scan(&id, &name, &price)

		if err != nil {

			fmt.Println("ERROR")

		}

		bookData = &BookData{

			Id:    id,

			Name:  name,

			Price: price,

		}

		sliceBook = append(sliceBook, *bookData)
	}

	return sliceBook

}