package impl

import (
	"fmt"
	"github.com/cho8833/duary_lambda/internal/auth/dto"
	"github.com/cho8833/duary_lambda/internal/auth/repository"
	"log"
)

type OIDCServiceImpl struct {
	repository repository.OIDCPublicKeyRepository
}

func NewOIDCService(repository *repository.OIDCPublicKeyRepository) *OIDCServiceImpl {
	return &OIDCServiceImpl{repository: *repository}
}

func (svc *OIDCServiceImpl) GetPublicKey(url string, provider string, kid string) (*dto.JWK, error) {
	// retrieve public jwk from DB
	certRes, err := svc.repository.FindPublicKeyInDB(provider)
	if err != nil {
		return nil, err
	}
	jwk := findMatchingKey(kid, certRes.Keys)

	if jwk == nil {
		log.Printf("there's no matching key with kid: %s, provider: %s, url: %s. Update Public Key", kid, provider, url)
		newCert, err := svc.repository.GetPublicJWK(url)
		if err != nil {
			return nil, err
		}

		err = svc.repository.SaveJWK(provider, newCert.Keys)
		if err != nil {
			return nil, err
		}
		jwk = findMatchingKey(kid, newCert.Keys)
	}
	// if there's no matching key in new cert, return err
	if jwk == nil {
		return nil, fmt.Errorf("no matching key with kid: %s, provider: %s, url: %s", kid, provider, url)
	}

	return jwk, nil
}

func findMatchingKey(kid string, jwks []dto.JWK) *dto.JWK {
	for _, val := range jwks {
		if kid == val.Kid {
			return &val
		}
	}
	return nil
}
