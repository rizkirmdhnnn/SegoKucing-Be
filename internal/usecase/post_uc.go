package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type PostUseCase struct {
	PostRepository *repository.PostRepository
	TagsRepo       *repository.TagRepository
	Validate       *validator.Validate
	Viper          *viper.Viper
	Log            *logrus.Logger
}

func NewPostUseCase(postRepository *repository.PostRepository, tagsRepo *repository.TagRepository, validate *validator.Validate, viper *viper.Viper, log *logrus.Logger) *PostUseCase {
	return &PostUseCase{
		PostRepository: postRepository,
		TagsRepo:       tagsRepo,
		Validate:       validate,
		Viper:          viper,
		Log:            log,
	}
}

func (c *PostUseCase) CreatePost(ctx context.Context, request *model.CreatePostRequest) (*model.CreatePostResponse, error) {
	c.Log.Info("Creating post")

	// Validate request
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Error("Validation failed: ", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Get user ID from context
	userId := ctx.Value("user_id").(int64)

	// Create post
	post := &entity.Posts{
		UserID:  userId,
		Content: request.PostInHtml,
	}

	// Call PostRepository to create post
	post, err = c.PostRepository.Create(post)
	if err != nil {
		c.Log.Errorf("Error creating post: %v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Failed to create post")
	}

	err = c.TagsRepo.AssignTagsToPost(ctx, post.ID, request.Tags)
	if err != nil {
		c.Log.Errorf("Error assigning tags to post: %v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Failed to assign tags to post")
	}

	return &model.CreatePostResponse{
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
	}, nil
}

// get all posts
func (c *PostUseCase) GetAllPosts(ctx context.Context, params *model.GetPostListParams) (*[]model.Post, *model.Meta, error) {
	// Get user ID from context
	userId := ctx.Value("user_id").(int64)
	posts, meta, err := c.PostRepository.GetAllPosts(ctx, userId, params)
	if err != nil {
		c.Log.Errorf("Error getting posts: %v", err)
		return nil, nil, fiber.NewError(fiber.StatusBadRequest, "Failed to get posts")
	}

	return posts, meta, nil
}
