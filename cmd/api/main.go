package main

import (
	_ "github.com/JoaoPauloLousada/go_rest_api_boilerplate/docs" // load generated docs
	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger" // fiber-swagger middleware
)

// @title			Fiber Boilerplate API
// @description	This is a sample server for a Fiber application.
// @version		1.0
// @host			localhost:3000
// @BasePath		/
func main() {
	app := fiber.New()

	// ExampleRoute godoc
	//	@Summary		Example route
	//	@Description	simple route that returns "Hello, World!"
	//	@Tags			example
	//	@Accept
	//	@Produce	json
	//	@Param
	//	@Success	200	{string}	"Hello, World!"
	//	@Failure	400	{object}	httputil.HTTPError
	//	@Failure	404	{object}	httputil.HTTPError
	//	@Failure	500	{object}	httputil.HTTPError
	//	@Router		/api/ [get]
	app.Get("/api/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Use swagger
	app.Get("/api/swagger/*", swagger.WrapHandler) // default

	// Start the Fiber app on port 3000
	app.Listen(":3000")
}
