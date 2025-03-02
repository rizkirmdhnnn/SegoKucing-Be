package controller

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
	"github.com/sirupsen/logrus"
)

type FriendController struct {
	friendUsecase *usecase.FriendUsecase
	log           *logrus.Logger
}

func NewFriendController(friendUsecase *usecase.FriendUsecase, log *logrus.Logger) *FriendController {
	return &FriendController{
		friendUsecase: friendUsecase,
		log:           log,
	}
}

func (f *FriendController) AddFriend(ctx *fiber.Ctx) error {
	ip := ctx.IP()
	userid := ctx.Locals("user_id").(int64)

	f.log.WithFields(logrus.Fields{
		"ip":      ip,
		"user_id": userid,
	}).Info("Received add friend request")

	request := new(model.AddFriendRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		f.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to parse request body")

		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
		return err
	}

	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)

	response, err := f.friendUsecase.AddFriend(newCtx, request)
	if err != nil {
		f.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to add friend")

		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
		return err
	}

	f.log.WithFields(logrus.Fields{
		"ip":      ip,
		"user_id": userid,
		"friend":  response,
	}).Info("Successfully added friend")

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": response.Message,
	})

	return nil
}

func (f *FriendController) GetFriendList(ctx *fiber.Ctx) error {
	user_id := ctx.Locals("user_id").(int64)
	ip := ctx.IP()

	limitStr := ctx.Query("limit", "10")
	offsetStr := ctx.Query("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		f.log.WithFields(logrus.Fields{
			"ip":    ip,
			"error": err,
		}).Error("Invalid limit value")
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid limit value",
		})
		return err
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		f.log.WithFields(logrus.Fields{
			"ip":    ip,
			"error": err,
		}).Error("Invalid offset value")
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid offset value",
		})
		return err
	}

	params := &model.GetFriendListParams{
		Limit:      limit,
		Offset:     offset,
		OrderBy:    ctx.Query("orderBy", "DESC"),
		SortBy:     ctx.Query("sortBy", "created_at"),
		Search:     ctx.Query("search", ""),
		OnlyFriend: ctx.Query("only_friend", "false") == "true",
	}

	f.log.WithFields(logrus.Fields{
		"ip":      ip,
		"user_id": user_id,
		"params":  params,
	}).Info("Fetching friend list")

	newCtx := context.WithValue(ctx.UserContext(), "user_id", user_id)

	response, err := f.friendUsecase.GetFriendList(newCtx, params)
	if err != nil {
		f.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to fetch friend list")
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
		return err
	}

	f.log.WithFields(logrus.Fields{
		"ip":      ip,
		"user_id": user_id,
		"count":   len(response.Friends),
	}).Info("Successfully fetched friend list")

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response.Friends,
		"meta": response.Meta,
	})

	return nil
}

func (f *FriendController) RemoveFriend(ctx *fiber.Ctx) error {
	user_id := ctx.Locals("user_id").(int64)
	ip := ctx.IP()

	f.log.WithFields(logrus.Fields{
		"ip":      ip,
		"user_id": user_id,
	}).Info("Received remove friend request")

	request := new(model.RemoveFriendRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		f.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to parse remove friend request body")
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
		return err
	}

	newCtx := context.WithValue(ctx.UserContext(), "user_id", user_id)

	response, err := f.friendUsecase.RemoveFriend(newCtx, request)
	if err != nil {
		f.log.WithFields(logrus.Fields{
			"ip":  ip,
			"err": err,
		}).Error("Failed to remove friend")
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
		return err
	}

	f.log.WithFields(logrus.Fields{
		"ip":      ip,
		"user_id": user_id,
		"message": response.Message,
	}).Info("Successfully removed friend")

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": response.Message,
	})

	return nil
}
