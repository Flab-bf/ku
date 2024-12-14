package model

import "time"

type Comment struct {
	Id        int       `gorm:"primary_key"`
	UserId    int       `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	CreateAt  time.Time `gorm:"type:datetime;"`
	UpdateAt  time.Time `gorm:"type:datetime;"`
	IsDeleted bool      `gorm:"default:false"`
	ParentId  *int
}
