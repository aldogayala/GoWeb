package handler

import (
	"net/http"
	"strconv"

	internal "github.com/aldogayala/GoWeb/internal/domain"
	productInternal "github.com/aldogayala/GoWeb/internal/domain/product"
	"github.com/gin-gonic/gin"
)

type Product struct {
	sv productInternal.Service
}

func NewProductHandler(sv productInternal.Service) *Product {
	return &Product{sv: sv}
}

func (pr *Product) GetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//request

		//process
		products, err := pr.sv.Get()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "All data of products",
			"data":    products,
		})

	}
}

func (pr *Product) GetProductByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//request
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		products, err := pr.sv.Get()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
		}

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

func (pr *Product) SearchByPrice() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		products, err := pr.sv.Get()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
		}

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

func (pr *Product) Create(name string, quantity int, codeValue string, isPublished bool, price float64) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func GetPong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong")
	}
}
