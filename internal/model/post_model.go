package model

import (
	"time"

	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
)

type CreatePostRequest struct {
	PostInHtml string `json:"postInHtml" validate:"required"`

	Tags []string `json:"tags" validate:"required"`
}

type CreatePostResponse struct {
	Content   string        `json:"content"`
	Tag       []entity.Tags `json:"tag,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
}

//TODO: Tag belum diimplementasi
