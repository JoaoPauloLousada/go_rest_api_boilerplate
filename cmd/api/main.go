package main

import (
	"fmt"
	"log"

	_ "github.com/JoaoPauloLousada/go_rest_api_boilerplate/docs" // load generated docs
	"github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/controllers"
	"github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/core/usecases"
	configuration "github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/infrastructure/config"
	persistence "github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/infrastructure/db"
	persistence_repositories "github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/infrastructure/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	fiberSwagger "github.com/swaggo/fiber-swagger" // fiber-swagger middleware
	// _ "github.com/swaggo/fiber-swagger/example/docs"
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

	db, err := persistence.ConnectDB(config)

	if err != nil {
		formattedError := fmt.Errorf("database connection error: %v", err)
		log.Fatal(formattedError)
	}

	app := fiber.New()

	SetupRoutes(app, db)

	// Use swagger
	app.Get("/api/swagger/*", fiberSwagger.WrapHandler) // default

	// Start the Fiber app on port 3000
	port := fmt.Sprintf(":%d", config.ApiPort)
	fmt.Printf("%s\u2714 Api is listening to the port:%d%s\n", "\033[32m", config.ApiPort, "\033[0m")
	app.Listen(port)
}

func SetupRoutes(app *fiber.App, db *sqlx.DB) {
	userRepository := persistence_repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUseCase(userRepository)
	userController := controllers.NewUserController(userUsecase)

	userGroup := app.Group("/api/v1/users")
	userGroup.Get("/", userController.GetAll)
	userGroup.Post("/", userController.Create)

	fmt.Printf("%s\u2714 Routes have been initialized%s\n", "\033[32m", "\033[0m")
}
