package controllers

import (
	domain "github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/core/entity"
	"github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/core/usecases"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	useCase *usecases.UserUseCase
}

func NewUserController(useCase *usecases.UserUseCase) *UserController {
	return &UserController{useCase: useCase}
}

// GetAll godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/users [get]
func (controller *UserController) GetAll(c *fiber.Ctx) error {
	controller.useCase.GetAll()
	users, err := controller.useCase.GetAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(users)
}

// Create godoc
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.User true "User object"
// @Success 201
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/users [post]
func (controller *UserController) Create(c *fiber.Ctx) error {
	var user domain.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := controller.useCase.Create(user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User created successfully"})
}
