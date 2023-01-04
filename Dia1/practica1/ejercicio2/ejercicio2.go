package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Nombre   string
	Apellido string
}

func (p Person) NombreCompleto() string {
	return fmt.Sprintf("Holi %s %s", p.Nombre, p.Apellido)
}

func main() {

	router := gin.Default()
	router.POST("/saludo", func(c *gin.Context) {
		var persona Person
		err := c.BindJSON(&persona)
		if err != nil {
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}

		c.String(http.StatusOK, persona.NombreCompleto())
	})
	router.Run()
}

/* Body sended:
{
    "Nombre":"Ange",
    "Apellido": "Marin"

}*/
