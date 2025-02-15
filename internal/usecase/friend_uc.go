package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/repository"
	"github.com/spf13/viper"
)

type FriendUsecase struct {
	friendRepo *repository.FriendRepository
	userRepo   *repository.UserRepository
	validator  *validator.Validate
	viper      *viper.Viper
}

func NewFriendUsecase(friendRepo *repository.FriendRepository, userRepo *repository.UserRepository, validator *validator.Validate, viper *viper.Viper) *FriendUsecase {
	return &FriendUsecase{
		friendRepo: friendRepo,
		userRepo:   userRepo,
		validator:  validator,
		viper:      viper,
	}
}

func (f *FriendUsecase) AddFriend(ctx context.Context, req *model.AddFriendRequest) (*model.AddFriendResponse, error) {
	if err := f.validator.Struct(req); err != nil {
		return nil, err
	}
	friend := req.ToEntity()
	friend.UserID = ctx.Value("user_id").(int64)

	// if friendID is self
	if friend.UserID == friend.FriendID {
		return nil, fiber.NewError(fiber.StatusBadRequest, "You can't add yourself as friend")
	}

	_, err := f.userRepo.GetUserById(int(friend.FriendID))
	if err != nil {
		return nil, err
	}

	// check if friend already exist
	_, err = f.friendRepo.GetFriendByUserIDAndFriendID(ctx, friend.UserID, friend.FriendID)
	if err == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Friend already exist")
	}

	err = f.friendRepo.AddFriend(ctx, friend)
	if err != nil {
		return nil, err
	}

	return &model.AddFriendResponse{
		Message: "Success",
	}, nil
}
