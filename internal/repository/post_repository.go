package repository

import (
	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		DB: db,
	}
}

func (r *PostRepository) Create(post *entity.Posts) (*entity.Posts, error) {
	err := r.DB.Create(post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepository) GetByID(id int64) (*entity.Posts, error) {
	var post entity.Posts
	err := r.DB.First(&post, id).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}
