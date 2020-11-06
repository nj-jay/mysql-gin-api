package main

import (

	"github.com/nj-jay/httpServer/routers"

	"log"

)

func main() {

	r := routers.Router()

	err := r.Run(":8080")

	if err != nil {

		log.Println("run error")

	}
}