package main

import (
	"go-rest-api/config"
	"go-rest-api/routes"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main () {

	app := fiber.New()
	
	routes.IndexRoutes(app.Group("/"))
	routes.PostRoutes(app.Group("/post"))

	err:= config.InitializeMongoDB(config.Config("MONGO_URI"), config.Config("MONGO_DATABASE"))

	if err != nil {
		defer config.DisconnectMongoDB()
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	port := config.Config("PORT")
	log.Fatal(app.Listen(port))
}