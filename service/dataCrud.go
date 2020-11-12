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

		err := rows.Scan(&bookData.Id, &bookData.Name, &bookData.Price)

		if err != nil {

			fmt.Println("ERROR")

		}

		sliceBook = append(sliceBook, *bookData)
	}

	return sliceBook

}

func DeleteSingleDataById(id string) Status {

	db, _ := database.Connect()

	defer db.Close()

	fmt.Println(id)

	_, err := db.Exec(database.DeleteSingleDataById, id)

	//fmt.Println(err)

	if err != nil {

		return 404

	} else {

		return 200
	}
}

func UpdateSingleDataById(id, name, price string) Status {

	db, _ := database.Connect()

	defer db.Close()

	fmt.Println(id, name, price)

	_, err := db.Exec(database.UpdateSingleDataById, name, price, id)

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