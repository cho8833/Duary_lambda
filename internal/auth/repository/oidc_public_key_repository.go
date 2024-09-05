package repository

import "github.com/cho8833/Duary/internal/auth/dto"

type OIDCPublicKeyRepository interface {
	FindPublicKeyInDB(provider string) (*dto.CertResponse, error)
	GetPublicJWK(url string) (*dto.CertResponse, error)
	SaveJWK(provider string, jwks []dto.JWK) error
}
