package repository

import "github.com/cho8833/CC-Calendar/internal/user/model"

type UserRepository interface {
	GetUser(userId int64) (*model.User, error)
}
