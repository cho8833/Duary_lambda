package service

import (
	"git-codecommit.ap-northeast-2.amazonaws.com/v1/repos/cc_calendar/internal/user/model"
	"git-codecommit.ap-northeast-2.amazonaws.com/v1/repos/cc_calendar/internal/user/repository"
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
