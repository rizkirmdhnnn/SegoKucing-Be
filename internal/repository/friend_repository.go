package repository

import (
	"context"

	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"gorm.io/gorm"
)

type FriendRepository struct {
	DB *gorm.DB
}

func NewFriendRepository(db *gorm.DB) *FriendRepository {
	return &FriendRepository{
		DB: db,
	}
}

func (f *FriendRepository) AddFriend(ctx context.Context, req *entity.Friends) error {
	err := f.DB.Create(req).Error
	if err != nil {
		return err
	}

	return nil
}

func (f *FriendRepository) GetFriendByUserIDAndFriendID(ctx context.Context, userID, friendID int64) (*entity.Friends, error) {
	var friend entity.Friends
	err := f.DB.Where("user_id = ? AND friend_id = ?", userID, friendID).First(&friend).Error
	if err != nil {
		return nil, err
	}

	return &friend, nil
}

func (f *FriendRepository) GetFriendCount(ctx context.Context, userID int64) (int, error) {
	var count int64
	err := f.DB.Model(&entity.Friends{}).Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (f *FriendRepository) GetFriendList(ctx context.Context, userID int64, limit, offset int) ([]entity.Friends, error) {
	var friends []entity.Friends
	err := f.DB.Where("user_id = ?", userID).Limit(limit).Offset(offset).Preload("Friend").Find(&friends).Error
	if err != nil {
		return nil, err
	}

	return friends, nil
}
