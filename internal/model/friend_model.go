package model

import "github.com/rizkirmdhnnn/segokucing-be/internal/entity"

type AddFriendRequest struct {
	UserId int64 `json:"userId" validate:"required"`
}

type AddFriendResponse struct {
	Message string `json:"message"`
}

func (a *AddFriendRequest) ToEntity() *entity.Friends {
	return &entity.Friends{
		FriendID: a.UserId,
	}
}
