package dao

import (
	"first/model"
	"first/utils"
	"time"
)

func IsRepeatUser(account string) (bool, error) {
	var count int64
	result := utils.DB.Model(&model.User{}).Where("account =?", account).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func CreateUser(user model.User) error {
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()
	result := utils.DB.Create(&user)
	return result.Error
}

func GetUserMessage(account string) (model.User, error) {
	var user model.User
	result := utils.DB.Where("account =?", account).First(&user)
	return user, result.Error
}