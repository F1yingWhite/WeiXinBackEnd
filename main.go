package main

import (
	"log"
	"weixin_backend/models"
	"weixin_backend/server"
)

func Init() {
	models.InitDB()
}

func main() {
	Init()
	api := server.InitRouter()

	err := api.Run(":8888")
	if err != nil {
		log.Panicln(err)
	}
}
