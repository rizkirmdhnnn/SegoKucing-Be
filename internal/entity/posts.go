package entity

import "time"

type Posts struct {
	ID        int64     `gorm:"column:id; primary_key; autoIncrement"`
	UserID    int64     `gorm:"column:user_id; not null"`
	Content   string    `gorm:"column:content; not null"`
	User      Users     `gorm:"foreignKey:UserID; references:ID"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:now()"`
}
