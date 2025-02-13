package entity

import "time"

type Users struct {
	ID        int64     `gorm:"column:id; primary_key; autoIncrement"`
	Name      string    `gorm:"column:name; not null"`
	Email     string    `gorm:"column:email; default:null; unique"`
	Phone     string    `gorm:"column:phone; default:null; unique"`
	Password  string    `gorm:"column:password; not null"`
	ImageUrl  string    `gorm:"column:image_url; default:null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:now()"`
}
