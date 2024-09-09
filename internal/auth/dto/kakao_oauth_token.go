package dto

import (
	"github.com/cho8833/duary_lambda/internal/auth/jwt_util"
	"time"
)

type KakaoOAuthToken struct {
	AccessToken           *string `json:"accessToken"`
	expiresAt             *time.Time
	refreshToken          *string
	refreshTokenExpiresAt *time.Time
	scopes                *[]string
	IdToken               *string `json:"idToken"`
}

type CertResponse struct {
	Keys []jwt_util.JWK `json:"keys"`
}
