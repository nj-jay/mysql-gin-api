package service

import (
	"errors"
	"fmt"
	"github.com/nj-jay/httpServer/global"
	"github.com/nj-jay/httpServer/middlewares"

	"github.com/nj-jay/httpServer/util"

	"github.com/nj-jay/httpServer/models"

	_ "github.com/nj-jay/httpServer/docs"

	//gs "github.com/swaggo/gin-swagger"
	//
	//"github.com/swaggo/gin-swagger/swaggerFiles"

)

type Login models.Login

//返回所有用户的登录信息
func QueryLogin() []models.Login {

	//rows, err := db.Query(database.QueryLogin)

	var logins []models.Login

	result := global.GMD_DB.Find(&logins)

	if result.Error != nil {

		return nil

	}

	return logins

}

//判断用户是否正确输入用户名和密码
func TrueLogin(username, password string) (string, error) {

	allLogin := QueryLogin()

	for _, value := range allLogin {

		if username == value.Username {

			//通过hash解密判断用户密码是否正确
			if util.HashDecrypt(value.Password, password) {

				//生成token
				tokenString, _ := middlewares.GenToken(value.Username, value.Password)

				return tokenString, nil

			} else {

				return "", errors.New("error")

			}
		}
	}

	return "", nil
}

//添加用户
func PostLogin(username, password string) Status {

	//password加密
	hashPwd := util.HashEncryption(password)

	fmt.Println(hashPwd)

	//stmt, _ := db.Prepare(database.PostUsername)

	//数据库保存hash加密后的password
	//_, err := stmt.Exec(username, hashPwd)


	login := &models.Login{

		Username: username,

		Password: hashPwd,

	}

	result := global.GMD_DB.Create(&login)

	if result.Error != nil {

		return 401

	} else {

		return 200

	}



}