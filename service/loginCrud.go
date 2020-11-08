package service

import (

	"fmt"

	"github.com/nj-jay/httpServer/database"

	"github.com/nj-jay/httpServer/models"

)

type Login models.Login

func QueryLogin() []Login {

	db, _ := database.Connect()

	defer db.Close()

	rows, err := db.Query(database.QueryLogin)

	sliceLogin := make([]Login, 0)

	login := new(Login)

	if err != nil {

		fmt.Println(err)

	}

	var username string

	var password string

	for rows.Next() {

		err := rows.Scan(&username, &password)

		if err != nil {

			fmt.Println("ERROR")

		}

		login = &Login{

			Username:	username,

			Password:	password,

		}

		sliceLogin = append(sliceLogin, *login)

	}

	return sliceLogin


}