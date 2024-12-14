package model

import "time"

type User struct {
	Id       int       `gorm:"primary_key"`
	Nickname string    `gorm:"type:varchar(225)"`
	Account  string    `gorm:"type:varchar(255);not null;unique"`
	Password string    `gorm:"type:varchar(255);not null"`
	CreateAt time.Time `gorm:"type:timestamp"`
	UpdateAt time.Time `gorm:"type:timestamp;table:user"`
}
