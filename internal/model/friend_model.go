package model

import (
	"time"

	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
)

type AddFriendRequest struct {
	UserId int64 `json:"userId" validate:"required"`
}

type AddFriendResponse struct {
	Message string `json:"message"`
}

type Friend struct {
	UserID      int64     `json:"userId"`
	Name        string    `json:"name"`
	ImageUrl    string    `json:"imageUrl"`
	FriendCount int       `json:"friendCount"`
	CreatedAt   time.Time `json:"createdAt"`
}

type GetFriendListResponse struct {
	Friends []Friend `json:"friends"`
	Meta    Meta     `json:"meta"`
}

type GetFriendListParams struct {
	Limit      int
	Offset     int
	SortBy     string
	OrderBy    string
	OnlyFriend bool
	Search     string
}

type RemoveFriendRequest struct {
	UserId int64 `json:"userId" validate:"required"`
}

type RemoveFriendResponse struct {
	Message string `json:"message"`
}

func (a *AddFriendRequest) ToEntity() *entity.Friends {
	return &entity.Friends{
		FriendID: a.UserId,
	}
}
