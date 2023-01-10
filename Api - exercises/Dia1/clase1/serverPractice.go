package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default() //Router por defecto. Logger: Crea un log en la consola y Recovery: Cada vez quqe hay un Panic busca la manera de recobrarse

	//Create a /ping route
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//Run the server
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

	//Enciende el servidor: go run serverPractice.go
	//Haz una petición: http://localhost:8080/ping

	/* FORMA LARGA DE CREAR UN SERVIDOR
	//Create http serve in the easiest way
	/*err := http.ListenAndServe(":9876", nil)
	if err != nil {
		panic(err)
	}*/ /*

		//Create a ping handler
		pingHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		//Create a router
		//Create a multiplexer (ruta que se la asigna a n hadler)
		router := http.NewServeMux()

		//Reister /ping handler
		router.Handle("/ping", pingHandler)

		//Create http serve in a manual way
		server := &http.Server{
			Addr:         ":8080",
			WriteTimeout: 10 * time.Second,
			Handler:      router,
		}

		//Start server
		server.ListenAndServe()
	*/
}
