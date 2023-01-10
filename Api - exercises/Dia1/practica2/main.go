package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var products []Product

// pkg
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New(message string, data interface{}) *Response {
	return &Response{Message: message, Data: data}
}

func main() {

	products = chargeProducts()

	router := gin.Default()

	//Create a /ping route
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})
	router.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "succed to get all products", "data": products})
	})
	router.GET("/product/:id", BuscarProducto) //Send a parameter

	router.GET("/product/search", BuscarProductoMayor)

	router.GET("/productsHighPrice", BuscarProductoMayorParam)

	//Run the server
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

	router.Run()
	fmt.Println(products[0])

}

// This helps you to get more than one parameter and return all the products with more price than
func BuscarProductoMayorParam(ctxt *gin.Context) {
	flag := true

	priceGt, _ := strconv.ParseFloat(ctxt.Query("priceGt"), 64)

	var productsHighPrice []Product
	for _, product := range products {
		if product.Price > priceGt {
			productsHighPrice = append(productsHighPrice, product)

			if flag {
				flag = false
			}

		}
	}

	ctxt.JSON(200, productsHighPrice)

}

// Este retornará los productos con precio mayor a priceGt
func BuscarProductoMayor(ctxt *gin.Context) {
	flag := true
	priceGt := 900.0
	var productsHighPrice []Product
	for _, product := range products {
		if product.Price > priceGt {
			productsHighPrice = append(productsHighPrice, product)

			if flag {
				flag = false
			}

		}
	}
	if flag {
		ctxt.JSON(400, gin.H{"message": "products not found", "data": nil}) //Es correcto poner primero que salió mal y después que está bien
		//ctxt.String(400, "Información del producto ¡No existe!")
	} else {
		ctxt.JSON(200, productsHighPrice)
	}

}

// Este handler verificará si la id que pasa el cliente existe en nuestra base de datos.
func BuscarProducto(ctxt *gin.Context) {
	flag := true
	for _, product := range products {
		idParam, err := strconv.Atoi(ctxt.Param("id"))
		if err != nil {
			fmt.Println(errors.New("algo está mal"))
		}
		if product.Id == idParam {
			ctxt.JSON(200, product)

			flag = false
			break
		}
	}
	if flag {
		ctxt.String(404, "Información del producto ¡No existe!")
	}

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

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}
