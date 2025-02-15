package controller

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
)

type PostController struct {
	postUC *usecase.PostUseCase
}

func NewPostController(PostUC *usecase.PostUseCase) *PostController {
	return &PostController{
		postUC: PostUC,
	}
}

// CreatePost handles the creation of a new post.
// It extracts the user ID from the context, parses the request body into a CreatePostRequest model,
// and calls the PostUseCase to create the post. It returns a JSON response with the result.
func (c *PostController) CreatePost(ctx *fiber.Ctx) error {
	userid := ctx.Locals("user_id").(int64)
	request := new(model.CreatePostRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Printf("Error creating post for user_id %d: %v", userid, err)
		return fiber.ErrBadRequest
	}

	// Create new context with user_id value
	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)

	// Call PostUseCase to create post
	response, err := c.postUC.CreatePost(newCtx, request)
	if err != nil {
		log.Printf("Error creating post for user_id %d: %v", userid, err)
		return err
	}

	// Return JSON response
	return ctx.JSON(
		fiber.Map{
			"message": "Post created successfully",
			"data":    response,
		},
	)
}
