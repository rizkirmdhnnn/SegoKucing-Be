package repository

import (
	"context"

	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"gorm.io/gorm"
)

type TagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{
		DB: db,
	}
}

func (r *TagRepository) AssignTagsToPost(ctx context.Context, postID int64, tags []string) error {
	// Start transaction
	tx := r.DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existingTags []entity.Tags

	// Find all tags that already exist
	if err := tx.Where("tag IN ?", tags).Find(&existingTags).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Create a map of existing tags
	existingTagMap := make(map[string]entity.Tags)
	for _, tag := range existingTags {
		existingTagMap[tag.Tag] = tag
	}

	// Find all tags that don't exist
	var newTags []entity.Tags
	for _, tagName := range tags {
		if _, exists := existingTagMap[tagName]; !exists {
			newTags = append(newTags, entity.Tags{Tag: tagName})
		}
	}

	// Create new tags
	if len(newTags) > 0 {
		if err := tx.Create(&newTags).Error; err != nil {
			tx.Rollback()
			return err
		}
		existingTags = append(existingTags, newTags...)
	}

	// Find post
	var post entity.Posts
	if err := tx.First(&post, postID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Assign tags to post
	if err := tx.Model(&post).Association("Tags").Append(existingTags); err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
