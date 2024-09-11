package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/duary_lambda/internal/couple/model"
	"github.com/cho8833/duary_lambda/internal/couple/repository"
	"github.com/cho8833/duary_lambda/internal/couple/service"
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
)

func createCoupleAPI(_ context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// init
	cacheClient := util.GetCacheClient()
	dynamoDBClient, err := cacheClient.GetDynamoDBClient()
	if err != nil {
		log.Printf(err.Error())
		return util.LambdaAppErrorResponse(util.InternalServerError{}), nil
	}
	repo := repository.NewCoupleRepository(dynamoDBClient)
	svc := service.NewCoupleService(repo)

	createCoupleReq := &model.CreateCoupleReq{}
	err = json.Unmarshal([]byte(req.Body), &createCoupleReq)
	if err != nil {
		log.Printf(err.Error())
	}

	result, svcError := svc.CreateCouple(createCoupleReq)
	if err != nil {
		return util.LambdaAppErrorResponse(svcError), nil
	}
	return util.LambdaResponseWithData(result), nil
}

func main() {
	lambda.Start(createCoupleAPI)
}
