package repository

import (
	"context"

	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
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

func (f *FriendRepository) GetFriendList(ctx context.Context, userID int64, params *model.GetFriendListParams) ([]entity.Friends, model.Meta, error) {
	var friends []entity.Friends
	var total int64

	query := f.DB.Model(&entity.Friends{}).Where("user_id = ?", userID)

	// OnlyFriend: Menampilkan hanya teman yang juga berteman dengan user
	if params.OnlyFriend {
		query = query.Joins("JOIN friends f2 ON friends.friend_id = f2.user_id AND f2.friend_id = friends.user_id")
	}

	// Search by name atau email
	if params.Search != "" {
		query = query.Joins("JOIN users ON friends.friend_id = users.id").
			Where("users.name ILIKE ? OR users.email ILIKE ?", "%"+params.Search+"%", "%"+params.Search+"%")
	}

	// Hitung total data sebelum limit & offset
	err := query.Count(&total).Error
	if err != nil {
		return nil, model.Meta{}, err
	}

	// Sorting berdasarkan friendCount menggunakan subquery
	if params.SortBy == "friendCount" {
		query = query.Select("friends.*, COALESCE((SELECT COUNT(*) FROM friends f WHERE f.user_id = friends.friend_id), 0) AS friend_count").
			Order(gorm.Expr("friend_count " + params.OrderBy))
	} else {
		query = query.Order(params.SortBy + " " + params.OrderBy)
	}

	// Pagination (limit & offset)
	err = query.Limit(params.Limit).
		Offset(params.Offset).
		Preload("Friend").
		Find(&friends).Error

	if err != nil {
		return nil, model.Meta{}, err
	}

	// Buat meta data response
	meta := model.Meta{
		Limit:  params.Limit,
		Offset: params.Offset,
		Total:  int(total),
	}

	return friends, meta, nil
}
