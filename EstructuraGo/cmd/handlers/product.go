package handlers

import (
	"encoding/json"
	"net/http"
	"rest/internal/domain"
	"rest/internal/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	sv product.Service
}

func NewProduct(sv product.Service) *Product { //Más que un producto nuevo es un controlador de productos, no?
	return &Product{sv: sv}
}

//Acá están los handlers

func (p *Product) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request

		// process
		prod, err := p.sv.Get()
		if err != nil {
			ctx.JSON(500, nil)
			return
		}

		// response
		ctx.JSON(200, prod)
	}
}
func (w *Product) Create() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// request
		var req domain.Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, nil)
			return
		}

		// process
		prod, err := w.sv.Create(req.Name, req.Quantity, req.Code_value, req.Is_published, req.Expiration, req.Price)
		if err != nil {
			ctx.JSON(500, nil)
			return
		}

		// response
		ctx.JSON(200, prod)
	}
}

func (p *Product) Update() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// request
		var request domain.Request
		var prod domain.Product
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		// process
		prod, err = p.sv.Update(id, request.Name, request.Quantity, request.Code_value, request.Is_published, request.Expiration, request.Price)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		// response
		ctx.JSON(200, prod)
	}
}

func (p *Product) UpdatePartial() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		//request

		// Obtener el identificador del producto a actualizar.
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		// Obtener el producto a actualizar.
		product, err := p.sv.GetByID(id)
		if err != nil {
			/*switch err {
			case products.ErrProductNotFound:
				ctx.JSON(http.StatusNotFound, nil)
			default:
				ctx.JSON(http.StatusInternalServerError, nil)
			}*/
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		err = json.NewDecoder(ctx.Request.Body).Decode(&product) //¿qué pasa acá? -> Coge el body y reemplaza esos campos en el product
		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		//process
		prod, err := p.sv.Update(id, product.Name, product.Quantity, product.Code_value, product.Is_published, product.Expiration, product.Price) //Acá no sería mejor enviar un product ¿?
		if err != nil {
			/*switch err {
			case products.ErrProductNotFound:
				ctx.JSON(http.StatusNotFound, nil)
			default:
				ctx.JSON(http.StatusInternalServerError, nil)
			}*/
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		//response
		// Retornar el producto actualizado.
		ctx.JSON(http.StatusOK, prod)
	}

}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//request
		//get id
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		//process
		err = p.sv.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}

		//response
		ctx.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
	}

}
