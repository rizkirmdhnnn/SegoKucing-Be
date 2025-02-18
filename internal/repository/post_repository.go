package repository

import (
	"context"
	"log"
	"time"

	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"github.com/rizkirmdhnnn/segokucing-be/internal/model"
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

func (r *PostRepository) GetAllPosts(ctx context.Context, userID int64, params *model.GetPostListParams) (*[]model.Post, *model.Meta, error) {
	var posts []entity.Posts
	var totalPosts int64

	// Log to get total time query
	start := time.Now()
	defer func() {
		log.Printf("Total time query: %v", time.Since(start))
	}()

	// Hitung total post untuk meta pagination
	if err := r.DB.Model(&entity.Posts{}).Where("user_id = ?", userID).Count(&totalPosts).Error; err != nil {
		return nil, nil, err
	}

	// Query dengan sorting dan pagination
	query := r.DB.Preload("User").
		Preload("Tags").
		Preload("Comments.User").
		Limit(params.Limit).
		Offset(params.Offset).
		Where("user_id = ?", userID)

	if err := query.Find(&posts).Error; err != nil {
		return nil, nil, err
	}

	// Mengambil semua user ID yang terlibat dalam post dan komentar
	var userIDs []int64
	for _, post := range posts {
		userIDs = append(userIDs, post.UserID)
		for _, comment := range post.Comments {
			userIDs = append(userIDs, comment.UserID)
		}
	}

	// Menghitung jumlah teman untuk setiap user dalam satu query
	var friendCounts []struct {
		UserID int64
		Count  int
	}
	if err := r.DB.Table("friends").
		Select("user_id, COUNT(*) as count").
		Where("user_id IN (?)", userIDs).
		Group("user_id").
		Scan(&friendCounts).Error; err != nil {
		return nil, nil, err
	}

	// Membuat map untuk menyimpan friend count
	friendCountMap := make(map[int64]int)
	for _, fc := range friendCounts {
		friendCountMap[fc.UserID] = fc.Count
	}

	var responsePosts []model.Post

	for _, post := range posts {
		postResponse := model.Post{
			PostId: post.ID,
			Post: model.PostDetail{
				PostInHtml: post.Content,
				Tags:       entity.ExtractTags(post.Tags),
				CreatedAt:  post.CreatedAt,
			},
			Creator: model.Creator{
				UserID:      post.User.ID,
				Name:        post.User.Name,
				ImageUrl:    post.User.ImageUrl,
				FriendCount: friendCountMap[post.UserID],
			},
		}

		for _, comment := range post.Comments {
			postResponse.Comments = append(postResponse.Comments, model.Comment{
				Comment: comment.Comment,
				Creator: model.Creator{
					UserID:      comment.User.ID,
					Name:        comment.User.Name,
					ImageUrl:    comment.User.ImageUrl,
					FriendCount: friendCountMap[comment.UserID],
				},
				CreatedAt: comment.CreatedAt,
			})
		}

		responsePosts = append(responsePosts, postResponse)
	}

	meta := model.Meta{
		Total:  int(totalPosts),
		Limit:  params.Limit,
		Offset: params.Offset,
	}

	return &responsePosts, &meta, nil
}
