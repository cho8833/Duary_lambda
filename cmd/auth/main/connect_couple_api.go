package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/duary_lambda/internal/auth"
	"github.com/cho8833/duary_lambda/internal/common"
	"github.com/cho8833/duary_lambda/internal/connectcouple"
	"github.com/cho8833/duary_lambda/internal/couple"
	"github.com/cho8833/duary_lambda/internal/member"
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
)

/*
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -tags lambda.norpc -o bootstrap cmd/auth/main/connect_couple_api.go && chmod 755 bootstrap && zip  build/package/connect_couple_api.zip bootstrap && rm bootstrap
*/
func connectCouple(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cacheClient := util.GetCacheClient()
	dynamodbClient, _ := cacheClient.GetDynamoDBClient()

	transaction := util.NewWriteTransaction(dynamodbClient)
	memberRepo := member.NewMemberRepository(dynamodbClient)
	coupleRepo := couple.NewCoupleRepository(dynamodbClient)
	sessionRepo := connectcouple.NewConnectCoupleRepository(dynamodbClient)

	memberSvc := member.NewMemberService(memberRepo)
	coupleSvc := couple.NewCoupleService(coupleRepo)
	sessionSvc := connectcouple.NewConnectCoupleService(sessionRepo)

	commonSvc := common.NewCommonService(memberSvc, coupleSvc, sessionSvc)

	loginMember := auth.FromRequestContext(request)

	connectCoupleReq := &common.ConnectCoupleReq{}
	err := json.Unmarshal([]byte(request.Body), &connectCoupleReq)
	if err != nil {
		log.Printf(err.Error())
		return util.LambdaAppErrorResponse(util.BadRequestError{}), nil
	}

	res, svcError := commonSvc.ConnectCouple(loginMember, connectCoupleReq, transaction)
	if svcError != nil {
		return util.LambdaAppErrorResponse(svcError), nil
	}
	return util.LambdaResponseWithData(res), nil
}

func main() {
	lambda.Start(connectCouple)
}
