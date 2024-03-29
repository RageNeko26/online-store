package main

import (
	"fmt"
	"os"

	"github.com/RageNeko26/online-store/controller"
	"github.com/RageNeko26/online-store/db"
	_ "github.com/RageNeko26/online-store/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Synapsis Online Store API
// @version 1.0
// @description An API Documentation Online Store App
// @contact_name RidhoGAPX
// @contact_email ridhogalihpambudi57@gmail.com
// @host ec2-3-85-238-209.compute-1.amazonaws.com:3000
// @BasePath /api/v1/

func main() {
	app := fiber.New()

	// Load up environment variable
	SECRET := os.Getenv("SECRET")

	// Initialize configuration
	q := db.NewDBConnection(&db.DBCredentials{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	})

	controller := controller.Setup(app, q, []byte(SECRET))
	// Initialize route
	controller.Routes()

	// Swagger
	controller.App.Get("/swagger/*", swagger.HandlerDefault)

	// Internal logging
	fmt.Println("Server is running")

	app.Listen(":3000")
}
