package main

import (
	"github.com/aldogayala/GoWeb/cmd/server/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	//Init Gin
	routes := gin.Default()

	//Group routes
	pingGR := routes.Group("/ping")
	pingGR.GET("/", handler.GetPong())

	productsGR := routes.Group("/products")
	productsGR.GET("/", handler.GetProducts())
	productsGR.GET("/:id", handler.GetProductByID())
	productsGR.GET("/search/:priceGt", handler.SearchByPrice())

	//Run Server port: 9090
	routes.Run(":9090")

}
