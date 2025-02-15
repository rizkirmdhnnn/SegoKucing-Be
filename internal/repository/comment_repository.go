package repository

import (
	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		DB: db,
	}
}

func (r *CommentRepository) Create(comment *entity.Comments) (*entity.Comments, error) {
	err := r.DB.Create(comment).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}
