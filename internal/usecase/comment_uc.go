package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/repository"
	"github.com/spf13/viper"
)

type CommentUseCase struct {
	commentRepo *repository.CommentRepository
	postRepo    *repository.PostRepository
	Validator   *validator.Validate
	Viper       *viper.Viper
}

func NewCommentUseCase(commentRepo *repository.CommentRepository, postRepo *repository.PostRepository, validate *validator.Validate, viper *viper.Viper) *CommentUseCase {
	return &CommentUseCase{
		commentRepo: commentRepo,
		postRepo:    postRepo,
		Validator:   validate,
		Viper:       viper,
	}
}

func (uc *CommentUseCase) CreateComment(ctx context.Context, comment *model.CreateCommentRequest) (*model.CreateCommentResponse, error) {
	err := uc.Validator.Struct(comment)
	if err != nil {
		return nil, err
	}

	// check if post exist
	_, err = uc.postRepo.GetByID(comment.PostID)
	if err != nil {
		return nil, err
	}

	// model to entity
	commentEntity := comment.ToEntity()
	commentEntity.UserID = ctx.Value("user_id").(int64)

	// create comment
	resp, err := uc.commentRepo.Create(commentEntity)
	if err != nil {
		return nil, err
	}

	// return comment
	return &model.CreateCommentResponse{
		ID:        resp.ID,
		Comment:   resp.Comment,
		PostID:    resp.PostID,
		CreatedAt: resp.CreatedAt,
	}, nil
}
