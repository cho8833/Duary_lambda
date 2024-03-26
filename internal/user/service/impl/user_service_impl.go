package impl

import (
	"github.com/cho8833/CC-Calendar/internal/user/model"
	"github.com/cho8833/CC-Calendar/internal/user/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: *userRepository}
}

func (svc UserServiceImpl) GetUser(userId int64) (*model.User, error) {
	user, err := svc.userRepository.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
