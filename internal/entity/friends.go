package entity

import "time"

type Friends struct {
	ID        int64     `gorm:"column:id; primary_key; autoIncrement"`
	UserID    int64     `gorm:"column:user_id; not null"`
	FriendID  int64     `gorm:"column:friend_id; not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:now()"`
	User      Users     `gorm:"foreignKey:UserID; references:ID"`
	Friend    Users     `gorm:"foreignKey:FriendID; references:ID"`
}
