package main

import (
	"net/http"
	"strconv"

	"github.com/aldogayala/GoWeb/internal"
	"github.com/gin-gonic/gin"
)

func main() {

	//Init Gin
	routes := gin.Default()

	//Group routes
	pingGR := routes.Group("/ping")
	pingGR.GET("/", GetPong())

	productsGR := routes.Group("/products")
	productsGR.GET("/", GetProducts())
	productsGR.GET("/:id", GetProductByID())
	productsGR.GET("/search/:priceGt", SearchByPrice())

	//Run Server port: 9090
	routes.Run(":9090")

}

func GetPong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong")
	}
}

func GetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products := internal.LoadData()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "All data of products",
			"data":    products,
		})

	}
}

func GetProductByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		products := internal.LoadData()

		var result internal.Product
		for _, value := range products {
			if value.Id == id {
				result = value
				break
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Data of products ByID",
			"data":    result,
		})
	}
}

func SearchByPrice() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		products := internal.LoadData()

		var result []internal.Product

		price, err := strconv.Atoi(ctx.Param("priceGt"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		for _, value := range products {
			if value.Price > float64(price) {
				result = append(result, value)
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "All data of parameter",
			"data":    result,
		})

	}

}
