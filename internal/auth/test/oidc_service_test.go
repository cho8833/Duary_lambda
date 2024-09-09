package test

import (
	"fmt"
	oidcRepository "github.com/cho8833/Duary/internal/auth/repository"
	repository "github.com/cho8833/Duary/internal/auth/repository/impl"
	oidcService "github.com/cho8833/Duary/internal/auth/service/impl"
	testUtil "github.com/cho8833/Duary/internal/test/util"
	"github.com/cho8833/Duary/internal/util"
	"testing"
)

func Test_updatePublicKey(t *testing.T) {
	cacheClient := util.GetCacheClient()
	httpClient, err := cacheClient.GetHttpClient()
	if err != nil {
		t.Fatalf(err.Error())
	}
	dynamodbClient := testUtil.CreateLocalDynamoDBClient()

	var repo oidcRepository.OIDCPublicKeyRepository = repository.NewOIDCPublicKeyRepository(httpClient, dynamodbClient)

	svc := oidcService.NewOIDCService(&repo)

	certRes, err := svc.GetPublicKey("https://kauth.kakao.com/.well-known/jwks.json", "kakao", "")
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Print(certRes)
}
