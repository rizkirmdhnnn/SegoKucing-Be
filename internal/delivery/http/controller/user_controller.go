package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
)

type UserController struct {
	userUC *usecase.UserUseCase
}

func NewUserController(userUC *usecase.UserUseCase) *UserController {
	return &UserController{
		userUC: userUC,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Printf("Error parsing request body in Register: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Failed to register user")
	}

	response, err := c.userUC.CreateUser(ctx.UserContext(), request)
	if err != nil {
		log.Printf("Error logging in user: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Failed to register user")
	}

	return ctx.JSON(
		fiber.Map{
			"message": "User registered successfully",
			"data":    response,
		},
	)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "Failed to login")
	}

	response, err := c.userUC.Login(ctx.UserContext(), request)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "Failed to login")
	}

	return ctx.JSON(
		fiber.Map{
			"message": "User logged in successfully",
			"data":    response,
		},
	)
}
