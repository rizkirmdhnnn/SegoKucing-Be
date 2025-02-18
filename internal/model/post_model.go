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

type GetPostListParams struct {
	Limit     int      `json:"limit"`
	Offset    int      `json:"offset"`
	Search    string   `json:"search"`
	SearchTag []string `json:"searchTag"`
}

type Creator struct {
	UserID      int64  `json:"userId"`
	Name        string `json:"name"`
	ImageUrl    string `json:"imageUrl"`
	FriendCount int    `json:"friendCount"`
}

type Comment struct {
	Comment   string    `json:"comment"`
	Creator   Creator   `json:"creator"`
	CreatedAt time.Time `json:"createdAt"`
}

type Post struct {
	PostId   int64      `json:"postId"`
	Post     PostDetail `json:"post"`
	Comments []Comment  `json:"comments"`
	Creator  Creator    `json:"creator"`
}

type PostDetail struct {
	PostInHtml string    `json:"postInHtml"`
	Tags       []string  `json:"tags"`
	CreatedAt  time.Time `json:"createdAt"`
}

type GetPostListResponse struct {
	Data []Post `json:"data"`
	Meta Meta   `json:"meta"`
}
