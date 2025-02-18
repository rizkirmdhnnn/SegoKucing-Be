package entity

type Tags struct {
	ID  int64  `gorm:"column:id; primary_key; autoIncrement"`
	Tag string `gorm:"column:tag; not null; uniqueIndex"`
}

func ExtractTags(tags []Tags) []string {
	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Tag)
	}
	return tagNames
}
