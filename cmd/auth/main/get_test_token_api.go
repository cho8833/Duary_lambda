package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/duary_lambda/internal/auth/jwtutil"
	"github.com/cho8833/duary_lambda/internal/member"
	"github.com/cho8833/duary_lambda/internal/util"
	"os"
)

/*
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -tags lambda.norpc -o bootstrap cmd/auth/main/get_test_token_api.go && chmod 755 bootstrap && zip  build/package/get_test_token_api.zip bootstrap && rm bootstrap
*/
func getTestToken(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	key := os.Getenv("secretKey")

	jwtUtil := jwtutil.Impl{}
	cacheClient := util.GetCacheClient()
	dynamodbClient, _ := cacheClient.GetDynamoDBClient()

	memberRepository := member.NewMemberRepository(dynamodbClient)
	findMember, err := memberRepository.FindBySocialIdAndProvider(3428835809, "kakao")
	if err != nil {
		return util.LambdaErrorResponse(err, 500), nil
	}

	token := jwtUtil.NewToken(jwtUtil.GenerateSubject(findMember), key)

	return util.LambdaResponseWithData(token), nil
}

func main() {
	lambda.Start(getTestToken)
}
