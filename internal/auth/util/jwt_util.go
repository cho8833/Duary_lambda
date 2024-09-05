package util

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/cho8833/Duary/internal/auth/dto"
	"github.com/golang-jwt/jwt/v5"
	"math/big"
	"time"
)

type ValidatingValue struct {
	iss   *string
	aud   *string
	nonce *string
	cert  *dto.CertResponse
}

func NewValidatingValue(iss *string, aud *string, nonce *string, cert *dto.CertResponse) *ValidatingValue {
	return &ValidatingValue{iss: iss, aud: aud, nonce: nonce, cert: cert}
}

func VerifyRS256(idToken string, value *ValidatingValue) error {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(idToken, claims, func(token *jwt.Token) (interface{}, error) {
		// verify with public key
		kid := token.Header["kid"].(string)
		key, err := findMatchingKey(value.cert, kid)
		if err != nil {
			return nil, err
		}
		n, err := decode(key.N)
		if err != nil {
			return nil, err
		}
		e, err := decode(key.E)
		if err != nil {
			return nil, err
		}
		pk := &rsa.PublicKey{
			N: new(big.Int).SetBytes(n),
			E: int(new(big.Int).SetBytes(e).Int64()),
		}
		return pk, nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// check issuer
		iss, err := claims.GetIssuer()
		if err != nil {
			return err
		}
		if value.iss != nil {
			if iss != *value.iss {
				return fmt.Errorf("issuer does not match")
			}
		}

		// check audience
		aud, err := claims.GetAudience()
		if err != nil {
			return err
		}
		if value.aud != nil {
			if aud[0] != *value.aud {
				return fmt.Errorf("audience does not match")
			}
		}

		// check expirationTime, must be before now
		exp, err := claims.GetExpirationTime()
		if err != nil {
			return err
		}
		if exp.Before(time.Now()) {
			return fmt.Errorf("expire time is not valid")
		}

		// check nonce, pair with frontend
		nonce := claims["nonce"].(string)
		if value.nonce != nil {
			if nonce != *value.nonce {
				return fmt.Errorf("nonce does not match")
			}
		}
	}

	return nil
}

func findMatchingKey(response *dto.CertResponse, kid string) (*dto.JWK, error) {
	if response == nil {
		return nil, fmt.Errorf("response is null")
	}
	for _, key := range response.Keys {
		if key.Kid == kid {
			return &key, nil
		}
	}
	return nil, fmt.Errorf("idToken: could not find matching cert keyId for the token")
}

func decode(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}
