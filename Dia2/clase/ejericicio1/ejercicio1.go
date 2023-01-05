package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" validate:"required"`
	Quantity     int     `json:"quantity" validate:"required"`
	Code_value   string  `json:"code_value" validate:"required"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
}

var products []Product
var id = 0

func main() {
	products = chargeProducts()
	id = getLastID()

	//Setting the default server
	server := gin.Default()

	//Create routes

	//adding new product
	server.POST("/products", addNewProduct)

	//getting last product added
	server.GET("/products/:id", getProduct)

	err := server.Run(":9090")
	if err != nil {
		fmt.Println(errors.New(err.Error()))
		fmt.Println("request error: ", err)
	}

	server.Run()

}

func getProduct(ctx *gin.Context) {
	flag := true

	for _, p := range products {
		idRequested, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request sended"})
		}
		if idRequested == p.Id {
			ctx.JSON(http.StatusAccepted, gin.H{"product": p})
			flag = false
			return
		}

	}

	if flag {
		ctx.String(404, "Información del producto ¡No existe!")
	}

}

func addNewProduct(ctx *gin.Context) {

	// request
	var newProduct Product
	err := ctx.ShouldBind(&newProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "types invalid"})
		//ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	if err := validate.Struct(&newProduct); err != nil {

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "structure invalid"})
		return
	}

	// process

	//validating code
	for _, p := range products {
		if p.Code_value == newProduct.Code_value {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "code invalid"})
			return
		}
	}

	//validating date
	stringDate := newProduct.Expiration

	// parse string date to golang time
	t, err := time.Parse("01/02/2006", stringDate)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expiration date is invalid"})
		return
	}

	fmt.Println("Time: ", t, "Type of t: ", reflect.ValueOf(t).Type())

	id = id + 1
	newProduct.Id = id
	products = append(products, newProduct)

	// response
	ctx.JSON(http.StatusCreated, gin.H{"message": "product added succeed"})
}

func chargeProducts() (products []Product) {

	// Open our jsonFile
	jsonFile, err := os.Open("products.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &products)

	return

}

func getLastID() (id int) {
	id = products[len(products)-1].Id
	return
}
