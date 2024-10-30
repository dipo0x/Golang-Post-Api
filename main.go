package main

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main () {

	app := fiber.New()
	
	routes.PostRoutes(app.Group("/post"))

	fmt.Println(config.Config("MONGO_URI"), config.Config("MONGO_DATABASE"))
	err:= config.InitializeMongoDB(config.Config("MONGO_URI"), config.Config("MONGO_DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	port := config.Config("PORT")
	log.Fatal(app.Listen(port))
}