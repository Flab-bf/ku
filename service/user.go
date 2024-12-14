package service

import (
	"first/dao"
	"first/model"
)

type Registeruser struct {
	Nickname string `json:"nickname"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type Loginuser struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func Register(user Registeruser) error {
	is, err := dao.IsRepeatUser(user.Account)
	if err != nil {
		return err
	}
	if is {
		return ErrUserNameRepeat
	}

	creUser := model.User{
		Nickname: user.Nickname,
		Account:  user.Account,
		Password: user.Password,
	}
	return dao.CreateUser(creUser)
}

func Login(user Loginuser) error {
	storeUser, err := dao.GetUserMessage(user.Account)
	if err != nil {
		return err
	}
	if storeUser.Password != user.Password {
		return ErrUserNotFound
	}
	return nil
}
