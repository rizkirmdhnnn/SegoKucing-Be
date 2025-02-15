package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
)

type FriendController struct {
	friendUsecase *usecase.FriendUsecase
}

func NewFriendController(friendUsecase *usecase.FriendUsecase) *FriendController {
	return &FriendController{
		friendUsecase: friendUsecase,
	}
}

func (f *FriendController) AddFriend(ctx *fiber.Ctx) error {
	userid := ctx.Locals("user_id").(int64)
	request := new(model.AddFriendRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
		return err
	}

	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)

	response, err := f.friendUsecase.AddFriend(newCtx, request)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
		return err
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": response.Message,
	})

	return nil
}
