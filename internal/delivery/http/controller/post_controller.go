package controller

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/usecase"
	"github.com/sirupsen/logrus"
)

type PostController struct {
	postUC *usecase.PostUseCase
	log    *logrus.Logger
}

func NewPostController(PostUC *usecase.PostUseCase, log *logrus.Logger) *PostController {
	return &PostController{
		postUC: PostUC,
		log:    log,
	}
}

// CreatePost handles the creation of a new post.
func (c *PostController) CreatePost(ctx *fiber.Ctx) error {
	ip := ctx.IP()
	c.log.WithFields(logrus.Fields{
		"ip": ip,
	}).Info("Received create post request")

	userid, ok := ctx.Locals("user_id").(int64)
	if !ok {
		c.log.Error("Failed to retrieve user_id from context")
		return fiber.ErrUnauthorized
	}

	request := new(model.CreatePostRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"user_id": userid,
			"error":   err.Error(),
		}).Error("Error parsing request body")
		return fiber.ErrBadRequest
	}

	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)
	response, err := c.postUC.CreatePost(newCtx, request)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"user_id": userid,
			"error":   err.Error(),
		}).Error("Error creating post")
		return err
	}

	c.log.WithFields(logrus.Fields{
		"user_id": userid,
		"post":    response,
	}).Info("Post created successfully")

	return ctx.JSON(fiber.Map{
		"message": "Post created successfully",
		"data":    response,
	})
}

// GetPostList handles the retrieval of a list of posts.
func (c *PostController) GetPostList(ctx *fiber.Ctx) error {
	userid, ok := ctx.Locals("user_id").(int64)
	if !ok {
		c.log.Error("Failed to retrieve user_id from context")
		return fiber.ErrUnauthorized
	}

	request := new(model.GetPostListParams)
	err := ctx.QueryParser(request)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error parsing query parameters")
		return fiber.ErrBadRequest
	}

	params := &model.GetPostListParams{
		Limit:     ctx.QueryInt("limit", 10),
		Offset:    ctx.QueryInt("offset", 0),
		Search:    ctx.Query("search", ""),
		SearchTag: strings.Split(ctx.Query("searchTag", ""), ","),
	}

	newCtx := context.WithValue(ctx.UserContext(), "user_id", userid)
	posts, meta, err := c.postUC.GetAllPosts(newCtx, params)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"user_id": userid,
			"error":   err.Error(),
		}).Error("Error getting post list")
		return err
	}

	c.log.WithFields(logrus.Fields{
		"user_id": userid,
		"total":   meta.Total,
	}).Info("Post list retrieved successfully")

	return ctx.JSON(fiber.Map{
		"data": posts,
		"meta": meta,
	})
}
