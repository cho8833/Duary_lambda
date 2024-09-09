package util

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cho8833/Duary/internal/auth/dto"
	"github.com/cho8833/Duary/internal/util"
	"github.com/golang-jwt/jwt/v5"
	"math/big"
	"time"
)

type ValidatingValue struct {
	Iss      string
	Aud      string
	Nonce    string
	Url      string
	Provider string
}

func VerifyRSA256(idToken string, value *ValidatingValue) error {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(idToken, claims, func(token *jwt.Token) (interface{}, error) {
		// verify with public key
		kid := token.Header["kid"].(string)
		jwk, err := getPublicKey(kid, value.Url, value.Provider)
		if err != nil {
			return nil, err
		}

		n, err := decode(jwk.N)
		if err != nil {
			return nil, err
		}
		e, err := decode(jwk.E)
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
		if value.Iss != "" {
			if iss != value.Iss {
				return fmt.Errorf("issuer does not match")
			}
		}

		// check audience
		aud, err := claims.GetAudience()
		if err != nil {
			return err
		}
		if value.Aud != "" {
			if aud[0] != value.Aud {
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

		// check Nonce, pair with frontend
		nonce := claims["Nonce"].(string)
		if value.Nonce != "" {
			if nonce != value.Nonce {
				return fmt.Errorf("Nonce does not match")
			}
		}
	}

	return nil
}

func decode(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}

func getPublicKey(kid string, url string, provider string) (*dto.JWK, error) {
	payload, err := json.Marshal(&dto.GetPublicKeyReq{
		Url:      url,
		Provider: provider,
		Kid:      kid,
	})
	if err != nil {
		return nil, err
	}
	client, err := util.GetCacheClient().GetLambdaClient()
	invokeInput := lambda.InvokeInput{
		FunctionName: aws.String("get_oidc_public_key"),
		LogType:      types.LogTypeTail,
		Payload:      payload,
	}
	invokeOutput, err := client.Invoke(context.TODO(), &invokeInput)
	if err != nil {
		return nil, err
	}
	if invokeOutput.StatusCode != 200 {
		return nil, fmt.Errorf("%+v", invokeOutput.Payload)
	}
	jwkResponse := &util.ServerResponse{}
	err = json.Unmarshal(invokeOutput.Payload, jwkResponse)
	if err != nil {
		return nil, err
	}
	if jwkResponse.Status != 200 {
		return nil, fmt.Errorf(*jwkResponse.Message)
	}
	return jwkResponse.Data.(*dto.JWK), nil
}
