package service

import "github.com/cho8833/CC-Calendar/internal/user/model"

type UserService interface {
	GetUser(userId int64) (*model.User, error)
}
