package main

import (

	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var (

	username = "mysql_nj_jay_com"

	password = "6435426njnj"

	ip = "47.95.1.229"

	port = "3306"

	database = "mysql_nj_jay_com"

)

type BookData struct {

	ID int `json:"id"`

	NAME string `json:"name"`

	PRICE float32 `json:"price"`

}

func main() {

	//gin框架默认配置
	r := gin.Default()

	//url路径
	r.GET("/api/v1/book", QueryDATA)

	//监听端口88
	_ = r.Run(":88")

}

func Connect() (*sql.DB, error) {

	path := strings.Join([]string{username, ":", password,
		"@(", ip, ":", port, ")/", database}, "")

	db, err := sql.Open("mysql", path)

	if err != nil {

		return db, err

	}

	if err = db.Ping(); err != nil {

		return db, err

	}

	return db, nil
}

func QueryDATA(c *gin.Context) {


	//连接数据库
	db, err := Connect()

	if err != nil {

		fmt.Println("数据库连接失败")

		//os.Exit(0)

	}

	//程序结束关闭
	defer db.Close()

	//mysql语句 选择数据表

	QueryData := `SELECT *FROM test;`

	//执行mysql语句

	rows, err := db.Query(QueryData)

	//创建结构体切片

	Slicebook := make([]BookData, 0)

	bookdata := new(BookData)

	if err != nil {

		fmt.Println(err)
	}

	var id int

	var name string

	var price float32

	//for循环读取数据表中的内容并赋值
	for rows.Next() {

		err := rows.Scan(&id, &name, &price)

		if err != nil {

			fmt.Println("ERROR")

		}

		bookdata = &BookData{

			ID:    id,
			NAME:  name,
			PRICE: price,

		}

		Slicebook = append(Slicebook, *bookdata)
	}

	//返回到网页
	c.JSON(200, Slicebook)

}