package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/duary_lambda/internal/auth/dto"
	"github.com/cho8833/duary_lambda/internal/auth/jwt_util"
	"github.com/cho8833/duary_lambda/internal/auth/service"
	"github.com/cho8833/duary_lambda/internal/member/repository"
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
)

func kakaoSignInAPI(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// init
	cacheClient := util.GetCacheClient()
	dynamoDBClient, err := cacheClient.GetDynamoDBClient()
	if err != nil {
		log.Printf(err.Error())
		return util.LambdaAppErrorResponse(util.InternalServerError{}), nil
	}
	memberRepository := repository.NewMemberRepository(dynamoDBClient)
	svc := service.NewKakaoAuthService(&jwt_util.JWTValidatorImpl{}, &jwt_util.JWTUtilImpl{}, memberRepository)

	// parse request
	kakaoToken := &dto.KakaoOAuthToken{}
	err = json.Unmarshal([]byte(request.Body), &kakaoToken)
	if err != nil {
		log.Printf(err.Error())
		return util.LambdaAppErrorResponse(util.BadRequestError{}), nil
	}

	// process
	result, svcError := svc.SignIn(kakaoToken)
	if svcError != nil {
		return util.LambdaAppErrorResponse(svcError), nil
	}
	return util.LambdaResponseWithData(result), nil
}

func main() {
	lambda.Start(kakaoSignInAPI)
}
