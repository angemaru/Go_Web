package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default() //Router por defecto. Logger: Crea un log en la consola y Recovery: Cada vez quqe hay un Panic busca la manera de recobrarse

	//Create a /ping route
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	//Run the server
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}
