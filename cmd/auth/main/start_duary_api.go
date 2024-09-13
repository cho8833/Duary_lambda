package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/duary_lambda/internal/common"
	"github.com/cho8833/duary_lambda/internal/couple"
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
)

func startDuaryAPI(_ context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// init
	cacheClient := util.GetCacheClient()
	dynamoDBClient, err := cacheClient.GetDynamoDBClient()
	if err != nil {
		log.Printf(err.Error())
		return util.LambdaAppErrorResponse(util.InternalServerError{}), nil
	}
	repo := couple.NewCoupleRepository(dynamoDBClient)
	svc := couple.NewCoupleService(repo)

	createCoupleReq := &common.InitDuaryInfoReq{}
	err = json.Unmarshal([]byte(req.Body), &createCoupleReq)
	if err != nil {
		log.Printf(err.Error())
		return util.LambdaAppErrorResponse(util.BadRequestError{}), nil
	}

	result, svcError := svc.CreateCouple(createCoupleReq)
	if err != nil {
		return util.LambdaAppErrorResponse(svcError), nil
	}
	return util.LambdaResponseWithData(result), nil
}

func main() {
	lambda.Start(startDuaryAPI)
}
