package models

//login信息结构体
type Login struct {

	Username string `json:"username" gorm:"primaryKey"`

	Password string `json:"password"`

}