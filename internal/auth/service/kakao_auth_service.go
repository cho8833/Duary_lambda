package service

import "github.com/cho8833/duary_lambda/internal/auth/dto"

type KakaoAuthService interface {
	SignIn(token dto.KakaoOAuthToken) (dto.SignInRes, error)
}
