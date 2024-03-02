package main

import (
	"fmt"
	"log"

	_ "github.com/JoaoPauloLousada/go_rest_api_boilerplate/docs" // load generated docs
	configuration "github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/infrastructure/config"
	persistence "github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/infrastructure/db"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger" // fiber-swagger middleware
)

// @title			Fiber Boilerplate API
// @description	This is a sample server for a Fiber application.
// @version		1.0
// @host			localhost:3000
// @BasePath		/
func main() {
	config, err := configuration.LoadConfig()

	if err != nil {
		formattedError := fmt.Errorf("config loading error: %v", err)
		log.Fatal(formattedError)
	}

	_, err = persistence.ConnectDB(config)

	if err != nil {
		formattedError := fmt.Errorf("database connection error: %v", err)
		log.Fatal(formattedError)
	}

	app := fiber.New()

	app.Get("/api/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Use swagger
	app.Get("/api/swagger/*", swagger.WrapHandler) // default

	// Start the Fiber app on port 3000
	port := fmt.Sprintf(":%d", config.ApiPort)
	fmt.Printf("%s\u2714 Api is listening to the port:%d%s\n", "\033[32m", config.ApiPort, "\033[0m")
	app.Listen(port)
}
