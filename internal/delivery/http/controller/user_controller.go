package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
)

type UserController struct {
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase) *UserController {
	return &UserController{
		UseCase: useCase,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Println(err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.CreateUser(ctx.UserContext(), request)
	if err != nil {
		log.Println(err)
		return err
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
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		log.Println(err)
		return err
	}

	return ctx.JSON(
		fiber.Map{
			"message": "User logged in successfully",
			"data":    response,
		},
	)
}
