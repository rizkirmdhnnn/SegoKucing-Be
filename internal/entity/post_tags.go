package entity

type PostTags struct {
	ID     int64 `gorm:"column:id; primary_key; autoIncrement"`
	PostID int64 `gorm:"column:posts_id; not null"`
	TagID  int64 `gorm:"column:tags_id; not null"`
	Post   Posts `gorm:"foreignKey:PostID; references:ID"`
	Tag    Tags  `gorm:"foreignKey:TagID; references:ID"`
}
