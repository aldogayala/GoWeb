package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/aldogayala/GoWeb/cmd/server/routes"
	domain "github.com/aldogayala/GoWeb/internal/domain"
	"github.com/gin-gonic/gin"
)

func main() {

	//Intances
	dataBase := LoadData()

	//Enviorement
	/*
		err := godotenv.Load("/.env")

		if err != nil {
			log.Fatal("Error al intentar cargar el archivo .env")
		}
	*/

	/*
		usuario := os.Getenv("MY_USER")
		password := os.Getenv("MY_PASS")

		fmt.Println("Usuario sacado de .env ", usuario)
		fmt.Println("Password sacado de .env ", password)
	*/

	//Init Gin
	engine := gin.Default()

	//Group routes
	router := routes.NewRouter(engine, &dataBase)
	router.SetProducts()

	//Run Server port: 9090
	if err := engine.Run(":9090"); err != nil {
		log.Fatal(err)
	}

}

func LoadData() (result []domain.Product) {
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

	var products []domain.Product

	json.Unmarshal([]byte(byteValue), &products)

	return products
}
