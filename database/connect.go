package database

import (

	"database/sql"

	"fmt"

	"os"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"

	"strings"

)

func Connect() (*sql.DB, error) {

	viper.SetConfigName("config")

	viper.SetConfigType("toml")

	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig()

	if err != nil {

		fmt.Println("read error")

		os.Exit(0)
	}

	username := viper.GetString("mysql.username")

	password := viper.GetString("mysql.password")

	ip := viper.GetString("mysql.ip")

	port := viper.GetString("mysql.port")

	database := viper.GetString("mysql.database")

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