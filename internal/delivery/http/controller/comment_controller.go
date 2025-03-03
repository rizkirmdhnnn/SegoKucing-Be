package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
	"github.com/sirupsen/logrus"
)

type CommentController struct {
	commentUC *usecase.CommentUseCase
	log       *logrus.Logger
}

func NewCommentController(commentUC *usecase.CommentUseCase, log *logrus.Logger) *CommentController {
	return &CommentController{
		commentUC: commentUC,
		log:       log,
	}
}

func (c *CommentController) CreateComment(ctx *fiber.Ctx) error {
	ip := ctx.IP()
	c.log.WithFields(logrus.Fields{
		"ip": ip,
	}).Info("Received create comment request")

	userid, ok := ctx.Locals("user_id").(int64)
	if !ok {
		c.log.Error("Failed to retrieve user_id from context")
		return fiber.ErrUnauthorized
	}

	request := new(model.CreateCommentRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"user_id": userid,
			"error":   err.Error(),
		}).Error("Error parsing request body")
		return fiber.ErrBadRequest
	}

	// Create new context with user_id value
	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)

	// Call CommentUseCase to create comment
	response, err := c.commentUC.CreateComment(newCtx, request)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"user_id": userid,
			"error":   err.Error(),
		}).Error("Error creating comment")
		return err
	}

	c.log.WithFields(logrus.Fields{
		"user_id": userid,
		"post":    response,
	}).Info("Comment created successfully")

	// Return JSON response
	return ctx.JSON(
		fiber.Map{
			"message": "Comment created successfully",
			"data":    response,
		},
	)
}
