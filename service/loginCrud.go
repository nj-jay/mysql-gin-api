package service

import (

	"fmt"

	"github.com/nj-jay/httpServer/util"

	"github.com/nj-jay/httpServer/database"

	"github.com/nj-jay/httpServer/models"
)

type Login models.Login

//返回所有用户的登录信息
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

//判断用户是否正确输入用户名和密码
func TrueLogin(username, password string) Status {

	var status int

	allLogin := QueryLogin()

	for _, value := range allLogin {

		if username == value.Username {

			//通过hash解密判断用户密码是否正确
			if util.HashDecrypt(value.Password, password) {

				status = 200

			}
		}
	}

	return Status(status)
}

//添加用户
func PostLogin(username, password string) Status {

	//password加密
	hashPwd := util.HashEncryption(password)

	fmt.Println(hashPwd)

	db, _ := database.Connect()

	stmt, _ := db.Prepare(database.PostUsername)

	//数据库保存hash加密后的password
	_, err := stmt.Exec(username, hashPwd)

	if err != nil {

		fmt.Println(err)

		return 404

	} else {

		return 200

	}
}