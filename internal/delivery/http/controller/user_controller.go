package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	userUC *usecase.UserUseCase
	log    *logrus.Logger
}

func NewUserController(userUC *usecase.UserUseCase, log *logrus.Logger) *UserController {
	return &UserController{
		userUC: userUC,
		log:    log,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	ip := ctx.IP()
	c.log.WithFields(logrus.Fields{
		"ip": ip,
	}).Info("Received register request")

	request := new(model.RegisterUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to parse request body")
		return fiber.NewError(fiber.StatusBadRequest, "Failed to register user")
	}

	response, err := c.userUC.CreateUser(ctx.UserContext(), request)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to create user")
		return fiber.NewError(fiber.StatusBadRequest, "Failed to register user")
	}

	c.log.WithFields(logrus.Fields{
		"ip":   ip,
		"user": response.Name,
	}).Info("User registered successfully")

	return ctx.JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    response,
	})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	ip := ctx.IP()
	c.log.WithFields(logrus.Fields{
		"ip": ip,
	}).Info("Received login request")

	request := new(model.LoginUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to parse request body")
		return fiber.NewError(fiber.StatusBadRequest, "Failed to login")
	}

	response, err := c.userUC.Login(ctx.UserContext(), request)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Login failed")
		return fiber.NewError(fiber.StatusBadRequest, "Failed to login")
	}

	c.log.WithFields(logrus.Fields{
		"ip":   ip,
		"user": response.Name,
	}).Info("User logged in successfully")

	return ctx.JSON(fiber.Map{
		"message": "User logged in successfully",
		"data":    response,
	})
}

func (c *UserController) LinkEmail(ctx *fiber.Ctx) error {
	ip := ctx.IP()
	userid := ctx.Locals("user_id").(int64)
	c.log.WithFields(logrus.Fields{
		"ip":      ip,
		"user_id": userid,
	}).Info("Received link email request")

	request := new(model.LinkEmailRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to parse request body")
		return fiber.NewError(fiber.StatusBadRequest, "Failed to link email")
	}

	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)
	if err := c.userUC.LinkEmail(newCtx, request); err != nil {
		c.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to link email")
		return err
	}

	c.log.WithFields(logrus.Fields{
		"ip":      ip,
		"user_id": userid,
	}).Info("Email linked successfully")

	return ctx.JSON(fiber.Map{
		"message": "Email linked successfully",
	})
}

func (c *UserController) LinkPhoneNumber(ctx *fiber.Ctx) error {
	ip := ctx.IP()
	userid := ctx.Locals("user_id").(int64)
	c.log.WithFields(logrus.Fields{
		"ip":      ip,
		"user_id": userid,
	}).Info("Received link phone number request")

	request := new(model.LinkPhoneRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to parse request body")
		return fiber.NewError(fiber.StatusBadRequest, "Failed to link phone number")
	}

	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)
	if err := c.userUC.LinkPhoneNumber(newCtx, request); err != nil {
		c.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to link phone number")
		return err
	}

	c.log.WithFields(logrus.Fields{
		"ip":      ip,
		"user_id": userid,
	}).Info("Phone number linked successfully")

	return ctx.JSON(fiber.Map{
		"message": "Phone Number linked successfully",
	})
}
