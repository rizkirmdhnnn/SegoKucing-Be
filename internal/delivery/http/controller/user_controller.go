package controller

import (
	"context"
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

// link email
func (c *UserController) LinkEmail(ctx *fiber.Ctx) error {
	userid := ctx.Locals("user_id").(int64)
	request := new(model.LinkEmailRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "Failed to link email")
	}

	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)

	err = c.userUC.LinkEmail(newCtx, request)
	if err != nil {
		log.Println(err)
		return err
	}

	return ctx.JSON(
		fiber.Map{
			"message": "Email linked successfully",
		},
	)
}

func (c *UserController) LinkPhoneNumber(ctx *fiber.Ctx) error {
	userid := ctx.Locals("user_id").(int64)
	request := new(model.LinkPhoneRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "Failed to link email")
	}

	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)

	err = c.userUC.LinkPhoneNumber(newCtx, request)
	if err != nil {
		log.Println(err)
		return err
	}

	return ctx.JSON(
		fiber.Map{
			"message": "Phone Number linked successfully",
		},
	)
}
