package service

import (
	"github.com/cho8833/duary_lambda/internal/auth/dto"
	errors "github.com/cho8833/duary_lambda/internal/util"
)

type KakaoAuthService interface {
	SignIn(token *dto.KakaoOAuthToken) (dto.SignInRes, errors.ApplicationError)
}
