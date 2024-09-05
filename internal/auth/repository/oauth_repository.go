package repository

import "github.com/cho8833/Duary/internal/auth/model"

type OAuthTokenRepository interface {
	FindOAuthBySocialIdAndProvider(socialId int64, provider string) (*model.OAuthToken, error)
	SaveOAuthToken(token *model.OAuthToken) error
}
