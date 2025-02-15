package model

import (
	"time"

	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
)

type CreateCommentRequest struct {
	PostID  int64  `json:"postId" validate:"required"`
	Comment string `json:"comment" validate:"required,min=2,max=500"`
}

type CreateCommentResponse struct {
	ID        int64     `json:"id"`
	PostID    int64     `json:"postId"`
	UserID    int64     `json:"userId"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *CreateCommentRequest) ToEntity() *entity.Comments {
	return &entity.Comments{
		PostID:  r.PostID,
		Comment: r.Comment,
	}
}
