package controller

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
)

type CommentController struct {
	commentUC *usecase.CommentUseCase
}

func NewCommentController(commentUC *usecase.CommentUseCase) *CommentController {
	return &CommentController{
		commentUC: commentUC,
	}
}

func (c *CommentController) CreateComment(ctx *fiber.Ctx) error {
	userid := ctx.Locals("user_id").(int64)
	request := new(model.CreateCommentRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Printf("Error creating comment: %v", err)
		return fiber.ErrBadRequest
	}

	// Create new context with user_id value
	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)

	// Call CommentUseCase to create comment
	response, err := c.commentUC.CreateComment(newCtx, request)
	if err != nil {
		log.Printf("Error creating comment: %v", err)
		return err
	}

	// Return JSON response
	return ctx.JSON(
		fiber.Map{
			"message": "Comment created successfully",
			"data":    response,
		},
	)
}
