package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func _main() {
	/*jsonAsMap := map[string]interface{}{
		"key1": "value1",
		"key2": true,
	}
	fmt.Printf("%v\n", jsonAsMap)
	fmt.Printf("Type of jsonAsMap%T\n", jsonAsMap)*/

	//Marshal -> Transformar un valor en json
	//var mapAsJson string
	/*mapAsJson, err := json.Marshal(jsonAsMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(mapAsJson)
	fmt.Println(string(mapAsJson)) //Debes castearlo a string porque el marshal lo convierte a un JSON en bytes
	*/

	//Marshal -> Transformar un valor en json con identación
	/*
		mapAsJson, err := json.MarshalIndent(jsonAsMap, "", " ")
		if err != nil {
			panic(err)
		} //Se recomienda solo usar el Marshal

		fmt.Printf("%s", mapAsJson)*/

	/*
		//Marshal 2.0
		//Create a product
		product := Product{
			name:  "Galletitas",
			Price: 60,
			//IsPublished: true,
			Secret: "Soy un súper secreto, no me leas",
		}

		//Convert to JSON
		productAsJson, err := json.Marshal((product))
		if err != nil {
			panic(err)
		}

		fmt.Println(string(productAsJson))

	*/

	/*
		// json.Unmarshal-> Transformar un json en una estructura/interfaz
		jsonString := `{"Name":"Galletitas","Price":60,"IsPublished":true}`

		//Convert json to structure
		var product Product
		err := json.Unmarshal([]byte(jsonString), &product) //Debes usar & porque debes apuntar al 'product'
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v", product) //fíjate en las los nombres, IsPublished no funciona xq está definido as is_published
	*/

	/*
		//Encoder
		encoder := json.NewEncoder(os.Stdout)

		product := Product{
			Name:  "Tablet",
			Price: 60,
			//IsPublished: true,
			Secret: "Soy un súper secreto, no me leas",
		}

		encoder.Encode(product)//Te escribe directamente en la consola
	*/

	//decoder
	const jsonStream = `
	{"name":"Tablet","price":60}
	{"name":"Compu","price":300}
	`
	streaming := strings.NewReader(jsonStream)
	decoder := json.NewDecoder(streaming)

	for {
		var product Product
		err := decoder.Decode(&product)
		if err == io.EOF { //No hay nada que leer (End Of File}9
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", product)
	}

}

// Define los nombres de los keys
type Product struct { //Si no está en mayúscula no se exporta
	Name        string  `json:"name"` //Si no está en mayuscula no se exporta, todo lo que se convierte debe ser PÚBLICO!
	Price       float64 `json:"price"`
	IsPublished bool    `json:"is_published,omitempty"` //si no esta definido simplemente no aparece :)
	Secret      string  `json:"-"`                      //Lo omite
}
