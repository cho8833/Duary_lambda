package service

import (
	"github.com/cho8833/kakao_login_lambda/model"
	"github.com/cho8833/kakao_login_lambda/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{userRepository: &repository.UserRepository{}}
}

func (svc UserService) GetUser(userId int64) (*model.User, error) {
	user, err := svc.userRepository.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
