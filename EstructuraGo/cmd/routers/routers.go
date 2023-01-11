package routers

import (
	"rest/cmd/handlers"
	"rest/internal/domain"
	"rest/internal/product"

	"github.com/gin-gonic/gin"
)

type Router struct {
	db *[]domain.Product
	en *gin.Engine // lo colocas porque este route va a trabajar sobre la instancia de este server
}

// recuerda que los metodos New son como constructores que nos permiten crear y manejar estas instancias
func NewRouter(en *gin.Engine, db *[]domain.Product) *Router {
	return &Router{en: en, db: db}
}

func (r *Router) SetRoutes() {
	r.SetProduct()
}

// Product
func (r *Router) SetProduct() {
	// instances
	rp := product.NewRepository(r.db, 3)
	sv := product.NewService(rp)
	h := handlers.NewProduct(sv)

	pd := r.en.Group("/product")
	pd.GET("", h.Get())
	pd.POST("", h.Create())
	pd.PUT("/:id", h.Update())
	pd.PATCH("/:id", h.UpdatePartial())
	pd.DELETE("/:id", h.Delete())
}
