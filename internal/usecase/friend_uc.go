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

func (f *FriendUsecase) GetFriendList(ctx context.Context, limit, offset int) (*model.GetFriendListResponse, error) {
	userID := ctx.Value("user_id").(int64)

	friends, err := f.friendRepo.GetFriendList(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	friendCount, err := f.friendRepo.GetFriendCount(ctx, userID)
	if err != nil {
		return nil, err
	}

	meta := model.Meta{
		Limit:  limit,
		Offset: offset,
		Total:  friendCount,
	}

	friendModel := make([]model.Friend, 0)
	for _, friend := range friends {
		if friend.Friend.ImageUrl == "" {
			friend.Friend.ImageUrl = "https://ui-avatars.com/api/?name=" + friend.Friend.Name
		}
		friendModel = append(friendModel, model.Friend{
			UserID:      friend.FriendID,
			Name:        friend.Friend.Name,
			ImageUrl:    friend.Friend.ImageUrl,
			FriendCount: friendCount,
			CreatedAt:   friend.CreatedAt,
		})
	}

	return &model.GetFriendListResponse{
		Friends: friendModel,
		Meta:    meta,
	}, nil
}
