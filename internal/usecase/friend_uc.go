package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
	"github.com/rizkirmdhnnn/segokucing-be/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type FriendUsecase struct {
	friendRepo *repository.FriendRepository
	userRepo   *repository.UserRepository
	validator  *validator.Validate
	viper      *viper.Viper
	log        *logrus.Logger
}

func NewFriendUsecase(friendRepo *repository.FriendRepository, userRepo *repository.UserRepository, validator *validator.Validate, viper *viper.Viper, log *logrus.Logger) *FriendUsecase {
	return &FriendUsecase{
		friendRepo: friendRepo,
		userRepo:   userRepo,
		validator:  validator,
		viper:      viper,
		log:        log,
	}
}

func (f *FriendUsecase) AddFriend(ctx context.Context, req *model.AddFriendRequest) (*model.AddFriendResponse, error) {
	f.log.Infof("Adding friend: %+v", req)

	if err := f.validator.Struct(req); err != nil {
		f.log.Warnf("Validation failed: %v", err)
		return nil, err
	}
	friend := req.ToEntity()
	friend.UserID = ctx.Value("user_id").(int64)

	if friend.UserID == friend.FriendID {
		f.log.Warn("User attempted to add themselves as a friend")
		return nil, fiber.NewError(fiber.StatusBadRequest, "You can't add yourself as friend")
	}

	_, err := f.userRepo.GetUserById(int(friend.FriendID))
	if err != nil {
		f.log.Errorf("Failed to get user by ID: %v", err)
		return nil, err
	}

	_, err = f.friendRepo.GetFriendByUserIDAndFriendID(ctx, friend.UserID, friend.FriendID)
	if err == nil {
		f.log.Warn("Friend already exists")
		return nil, fiber.NewError(fiber.StatusBadRequest, "Friend already exist")
	}

	err = f.friendRepo.AddFriend(ctx, friend)
	if err != nil {
		f.log.Errorf("Failed to add friend: %v", err)
		return nil, err
	}

	f.log.Info("Friend added successfully")
	return &model.AddFriendResponse{
		Message: "Success",
	}, nil
}

func (f *FriendUsecase) GetFriendList(ctx context.Context, params *model.GetFriendListParams) (*model.GetFriendListResponse, error) {
	f.log.Infof("Fetching friend list with params: %+v", params)
	userID := ctx.Value("user_id").(int64)

	if params.SortBy != "created_at" && params.SortBy != "friendCount" {
		f.log.Warn("Invalid sort parameter")
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid order by")
	}

	friends, meta, err := f.friendRepo.GetFriendList(ctx, userID, params)
	if err != nil {
		f.log.Errorf("Failed to fetch friend list: %v", err)
		return nil, err
	}

	friendCount, err := f.friendRepo.GetFriendCount(ctx, userID)
	if err != nil {
		f.log.Errorf("Failed to fetch friend count: %v", err)
		return nil, err
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

	f.log.Info("Friend list fetched successfully")
	return &model.GetFriendListResponse{
		Friends: friendModel,
		Meta:    meta,
	}, nil
}

func (f *FriendUsecase) RemoveFriend(ctx context.Context, friend *model.RemoveFriendRequest) (*model.RemoveFriendResponse, error) {
	f.log.Infof("Removing friend: %+v", friend)
	userID := ctx.Value("user_id").(int64)

	_, err := f.userRepo.GetUserById(int(friend.UserId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			f.log.Warn("User not found")
			return nil, fiber.NewError(fiber.StatusBadRequest, "User not found")
		}
	}

	_, err = f.friendRepo.GetFriendByUserIDAndFriendID(ctx, userID, friend.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			f.log.Warn("Friend not found")
			return nil, fiber.NewError(fiber.StatusBadRequest, "Friend not found")
		}
	}

	err = f.friendRepo.RemoveFriend(ctx, userID, friend.UserId)
	if err != nil {
		f.log.Errorf("Failed to remove friend: %v", err)
		return nil, err
	}

	f.log.Info("Friend removed successfully")
	return &model.RemoveFriendResponse{
		Message: "Success",
	}, nil
}
