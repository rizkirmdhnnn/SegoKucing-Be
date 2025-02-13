package entity

import "time"

type Comments struct {
	ID        int64     `gorm:"column:id; primary_key; autoIncrement"`
	PostID    int64     `gorm:"column:post_id; not null"`
	UserID    int64     `gorm:"column:user_id; not null"`
	Comment   string    `gorm:"column:comment; not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:now()"`
	Post      Posts     `gorm:"foreignKey:PostID; references:ID"`
	User      Users     `gorm:"foreignKey:UserID; references:ID"`
}
