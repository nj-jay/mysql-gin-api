package models

//书籍信息结构体
type BookData struct {

	Id uint `json:"id" gorm:"primaryKey"`

	Name string `json:"name"`

	Types string `json:"types"`

	Author string `json:"author"`

	Price float32 `json:"price"`

	AddTime string `json:"addTime"`

}
