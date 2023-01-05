package internal

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var (
	products []Product
)

func LoadData() (result []Product) {
	jsonFile, err := os.Open("./products.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Open successfully")

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(byteValue), &products)

	return products
}
