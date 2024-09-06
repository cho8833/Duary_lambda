package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	oidcRepository "github.com/cho8833/Duary/internal/auth/repository"
	repository "github.com/cho8833/Duary/internal/auth/repository/impl"
	authService "github.com/cho8833/Duary/internal/auth/service/impl"
	"github.com/cho8833/Duary/internal/util"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// load client
	cacheClient := util.GetCacheClient()
	httpClient, err := cacheClient.GetHttpClient()
	if err != nil {
		return util.CreateLambdaResponse(util.ResponseFromError(err, 500)), nil
	}
	dynamodbClient, err := cacheClient.GetDynamoDBClient()
	if err != nil {
		return util.CreateLambdaResponse(util.ResponseFromError(err, 500)), nil
	}

	// init service
	var repo oidcRepository.OIDCPublicKeyRepository = repository.NewOIDCPublicKeyRepository(httpClient, dynamodbClient)
	svc := authService.NewOIDCService(&repo)

	url := request.QueryStringParameters["url"]
	provider := request.QueryStringParameters["provider"]
	kid := request.QueryStringParameters["kid"]
	res, err := svc.GetPublicKey(url, provider, kid)
	if err != nil {
		return util.CreateLambdaResponse(util.ResponseFromError(err, 400)), nil
	}
	return util.CreateLambdaResponse(util.ResponseWithData(res)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
