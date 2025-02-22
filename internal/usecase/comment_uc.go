package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/repository"
	"github.com/spf13/viper"
)

type CommentUseCase struct {
	commentRepo *repository.CommentRepository
	friendRepo  *repository.FriendRepository
	postRepo    *repository.PostRepository
	Validator   *validator.Validate
	Viper       *viper.Viper
}

func NewCommentUseCase(commentRepo *repository.CommentRepository, friendRepo *repository.FriendRepository, postRepo *repository.PostRepository, validate *validator.Validate, viper *viper.Viper) *CommentUseCase {
	return &CommentUseCase{
		commentRepo: commentRepo,
		friendRepo:  friendRepo,
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

	//if user is the post owner
	if commentEntity.UserID != commentEntity.PostID {
		// check if user is friend
		isFriend, err := uc.friendRepo.IsFriend(ctx, commentEntity.UserID, commentEntity.PostID)
		if err != nil {
			return nil, err
		}

		if !isFriend {
			return nil, fiber.NewError(fiber.ErrBadRequest.Code, "You are not friend with the post owner")
		}
	}

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
