package models

type BookData struct {

	Id int `json:"id"`

	Name string `json:"name"`

	Price float32 `json:"price"`

}

//func NewBookData() *BookData {
//
//	return &BookData{
//
//		Id:    0,
//
//		Name:  "",
//
//		Price: 0,
//	}
//
//}