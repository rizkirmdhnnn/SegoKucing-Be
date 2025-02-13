package entity

type Tags struct {
	ID  int64  `gorm:"column:id; primary_key; autoIncrement"`
	Tag string `gorm:"column:tag; not null"`
}
