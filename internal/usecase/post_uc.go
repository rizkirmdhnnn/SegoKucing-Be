package usecase

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/repository"
	"github.com/spf13/viper"
)

type PostUseCase struct {
	PostRepository *repository.PostRepository
	TagsRepo       *repository.TagRepository
	Validate       *validator.Validate
	Viper          *viper.Viper
}

func NewPostUseCase(postRepository *repository.PostRepository, tagsRepo *repository.TagRepository, validate *validator.Validate, viper *viper.Viper) *PostUseCase {
	return &PostUseCase{
		PostRepository: postRepository,
		TagsRepo:       tagsRepo,
		Validate:       validate,
		Viper:          viper,
	}
}

func (c *PostUseCase) CreatePost(ctx context.Context, request *model.CreatePostRequest) (*model.CreatePostResponse, error) {
	// Validate request
	err := c.Validate.Struct(request)
	if err != nil {
		log.Println(err)
		return nil, fiber.ErrBadRequest
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
		log.Println(err)
		return nil, fiber.ErrInternalServerError
	}

	err = c.TagsRepo.AssignTagsToPost(ctx, post.ID, request.Tags)
	if err != nil {
		return nil, err
	}

	return &model.CreatePostResponse{
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
	}, nil
}
