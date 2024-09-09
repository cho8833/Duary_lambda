package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/duary_lambda/internal/auth/dto"
	oidcRepository "github.com/cho8833/duary_lambda/internal/auth/repository"
	repository "github.com/cho8833/duary_lambda/internal/auth/repository/impl"
	authService "github.com/cho8833/duary_lambda/internal/auth/service/impl"
	"github.com/cho8833/duary_lambda/internal/util"
)

func HandleRequest(ctx context.Context, request *dto.GetPublicKeyReq) (*util.ServerResponse[any], error) {
	// check Req
	if request.Url == "" || request.Provider == "" || request.Kid == "" {
		return util.ResponseFromError(fmt.Errorf("Bad Request"), 400), nil
	}

	// load client
	cacheClient := util.GetCacheClient()
	httpClient, err := cacheClient.GetHttpClient()
	if err != nil {
		return util.ResponseFromError(err, 500), nil
	}
	dynamodbClient, err := cacheClient.GetDynamoDBClient()
	if err != nil {
		return util.ResponseFromError(err, 500), nil
	}

	// init service
	var repo oidcRepository.OIDCPublicKeyRepository = repository.NewOIDCPublicKeyRepository(httpClient, dynamodbClient)
	svc := authService.NewOIDCService(&repo)

	res, err := svc.GetPublicKey(request.Url, request.Provider, request.Kid)
	if err != nil {
		return util.ResponseFromError(err, 400), nil
	}
	return util.ResponseWithData(res), nil
}

func main() {
	lambda.Start(HandleRequest)
}
