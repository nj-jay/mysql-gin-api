package models

//响应
type Error struct {

	Code int `json:"code"`

	Message string `json:"message"`

}

//简单工厂模式
func NewError() *Error {

	return &Error{

		Code:    404,

		Message: "not found",

	}
}

func NewSucceedDelete() *Error {

	return &Error{

		Code:	200,

		Message:	"successfully delete singleData",
	}
}

func NewErrorDelete() *Error {

	return &Error{

		Code:    404,

		Message: "operate error",

	}
}

