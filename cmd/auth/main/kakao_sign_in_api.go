package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/duary_lambda/internal/auth/dto"
	"github.com/cho8833/duary_lambda/internal/auth/jwt_util"
	impl2 "github.com/cho8833/duary_lambda/internal/auth/service/impl"
	"github.com/cho8833/duary_lambda/internal/member/repository/impl"
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// init
	cacheClient := util.GetCacheClient()
	dynamoDBClient, err := cacheClient.GetDynamoDBClient()
	if err != nil {
		log.Printf(err.Error())
		return util.LambdaErrorResponse(fmt.Errorf("알 수 없는 서버 오류"), 500), nil
	}
	memberRepository := impl.NewMemberRepository(dynamoDBClient)
	svc := impl2.NewKakaoAuthService(&jwt_util.JWTValidatorImpl{}, &jwt_util.JWTUtilImpl{}, memberRepository)

	// parse request
	kakaoToken := &dto.KakaoOAuthToken{}
	err = json.Unmarshal([]byte(request.Body), &kakaoToken)
	if err != nil {
		log.Printf(err.Error())
		return util.LambdaErrorResponse(fmt.Errorf("잘못된 요청"), 400), nil
	}

	// process
	result, svcError := svc.SignIn(kakaoToken)
	if err != nil {
		return util.LambdaAppErrorResponse(svcError), nil
	}
	return util.LambdaResponseWithData(result), nil
}

func main() {
	lambda.Start(handleRequest)
}
