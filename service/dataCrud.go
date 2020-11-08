package service

import (

	"fmt"

	"github.com/nj-jay/httpServer/database"

	"github.com/nj-jay/httpServer/models"
)

type BookData models.BookData

type Status int

func QueryData() []BookData {

	db, _ := database.Connect()

	defer db.Close()

	rows, err := db.Query(database.QueryAllData)

	sliceBook := make([]BookData, 0)

	bookData := new(BookData)

	if err != nil {

		fmt.Println(err)
	}

	for rows.Next() {

		err := rows.Scan(&bookData.Id, &bookData.Name, &bookData.Price)

		if err != nil {

			fmt.Println("ERROR")

		}

		sliceBook = append(sliceBook, *bookData)
	}

	return sliceBook

}

func QuerySingleDataById(id string) (*BookData, Status) {

	db, _ := database.Connect()

	defer db.Close()

	bookData := new(BookData)

	rows := db.QueryRow(database.QuerySingleDataById, id)

	err :=  rows.Scan(&bookData.Id, &bookData.Name, &bookData.Price)

	if err != nil {

		return bookData, 404

	} else {

		return bookData, 200
	}
}

func DeleteSingleDataById(id string) Status {

	db, _ := database.Connect()

	defer db.Close()

	stmt, _ := db.Prepare(database.DeleteSingleDataById)

	_, err := stmt.Exec(id)

	if err != nil {

		return 404

	} else {

		return 200
	}
}

func UpdateSingleDataById(id, price string) Status {

	db, _ := database.Connect()

	defer db.Close()

	stmt, _ := db.Prepare(database.UpdateSingleDataById)

	_, err := stmt.Exec(price, id)

	if err != nil {

		return 404

	} else {

		return 200
	}
}


func PostSingleData(name, price string) Status {

	db, _ := database.Connect()

	defer db.Close()

	stmt, _ := db.Prepare(database.PostSingleData)

	_, err := stmt.Exec(name, price)

	if err != nil {

		return  404

	} else {

		return 200
	}

}