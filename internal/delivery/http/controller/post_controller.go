package controller

import (
	"context"
	"log"
	"strings"

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

// GetPostList handles the retrieval of a list of posts.
// It parses the query parameters into a GetPostListParams model and calls the PostUseCase to get the list of posts.
// It returns a JSON response with the result.
func (c *PostController) GetPostList(ctx *fiber.Ctx) error {
	userid := ctx.Locals("user_id").(int64)

	request := new(model.GetPostListParams)
	err := ctx.QueryParser(request)
	if err != nil {
		log.Printf("Error getting post list: %v", err)
		return fiber.ErrBadRequest
	}

	params := &model.GetPostListParams{
		Limit:     ctx.QueryInt("limit", 10),
		Offset:    ctx.QueryInt("offset", 0),
		Search:    ctx.Query("search", ""),
		SearchTag: strings.Split(ctx.Query("searchTag", ""), ","),
	}

	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)

	// Call PostUseCase to get all posts
	posts, meta, err := c.postUC.GetAllPosts(newCtx, params)
	if err != nil {
		log.Printf("Error getting post list: %v", err)
		return err
	}

	// Return JSON response
	return ctx.JSON(
		fiber.Map{
			"data": posts,
			"meta": meta,
		},
	)
}
