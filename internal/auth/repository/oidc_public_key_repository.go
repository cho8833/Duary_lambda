package repository

import (
	"github.com/cho8833/duary_lambda/internal/auth/dto"
	"github.com/cho8833/duary_lambda/internal/auth/jwt_util"
)

type OIDCPublicKeyRepository interface {
	FindPublicKeyInDB(provider string) (*dto.CertResponse, error)
	GetPublicJWK(url string) (*dto.CertResponse, error)
	SaveJWK(provider string, jwks []jwt_util.JWK) error
}
