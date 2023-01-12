package routes

import (
	handler "github.com/aldogayala/GoWeb/cmd/server/handler"
	domain "github.com/aldogayala/GoWeb/internal/domain"
	"github.com/aldogayala/GoWeb/internal/domain/product"
	"github.com/gin-gonic/gin"
)

type Router struct {
	eng *gin.Engine
	db  *[]domain.Product
}

func NewRouter(eng *gin.Engine, db *[]domain.Product) *Router {
	return &Router{eng: eng, db: db}
}

func SetRoutes(r *Router) {
	r.SetProducts()
}

func (r *Router) SetProducts() {

	repositoryProduct := product.NewRepository(r.db)
	serviceProduct := product.NewService(repositoryProduct)
	handlerProduct := handler.NewProductHandler(serviceProduct)

	productsGR := r.eng.Group("/products")
	productsGR.GET("/", handlerProduct.GetProducts())
	/*
		productsGR.GET("/:id", handler.GetProductByID())
		productsGR.GET("/search/:priceGt", handler.SearchByPrice())
	*/
}
