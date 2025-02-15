package model

import (
	"time"

	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
)

type CreatePostRequest struct {
	PostInHtml string   `json:"postInHtml" validate:"required,min=2,max=500"`
	Tags       []string `json:"tags" validate:"required"`
}

type CreatePostResponse struct {
	Content   string        `json:"content"`
	Tag       []entity.Tags `json:"tag,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
}
