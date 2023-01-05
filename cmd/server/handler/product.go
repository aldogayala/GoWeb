package handler

import (
	"net/http"
	"strconv"

	internal "github.com/aldogayala/GoWeb/internal/domain"
	"github.com/gin-gonic/gin"
)

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

		price, err := strconv.ParseFloat(ctx.Param("priceGt"), 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		for _, value := range products {
			if value.Price > price {
				result = append(result, value)
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "All data of parameter",
			"data":    result,
		})

	}

}

func GetPong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong")
	}
}
