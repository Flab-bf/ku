package service

import (
	"first/dao"
	"first/model"
)

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func WriteComment(req Comment) error {
	comment := model.Comment{
		UserId:  req.Id,
		Content: req.Content,
	}
	return dao.WriteComment(comment)
}

func DeleteComments(id int) error {

	return dao.DeleteComment(id)
}

func ReadAll() ([]model.Comment, error) {
	return dao.GetAllComments()
}
