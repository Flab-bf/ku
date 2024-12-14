package dao

import (
	"first/model"
	"first/utils"
	"time"
)

func WriteComment(comment model.Comment) error {
	comment.CreateAt = time.Now()
	comment.UpdateAt = time.Now()
	result := utils.DB.Create(&comment)
	return result.Error
}

func DeleteComment(id int) error {
	result := utils.DB.Model(&model.Comment{}).Where("user_id = ?", id).Update("is_deleted", true)
	return result.Error
}

func GetAllComments() ([]model.Comment, error) {
	var comments []model.Comment
	result := utils.DB.Where("is_deleted=?", false).Find(&comments)
	return comments, result.Error
}
