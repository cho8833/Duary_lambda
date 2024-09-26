package connectcouple

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

type ApiGatewayRepository interface {
	PostToConnect(url string, connectionId *string, data any) error
}

type ApiGatewayRepositoryImpl struct {
}

func NewApiGatewayRepository() *ApiGatewayRepositoryImpl {
	return &ApiGatewayRepositoryImpl{}
}

func (repository *ApiGatewayRepositoryImpl) PostToConnect(url string, connectionId *string, data any) error {
	// get client
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	cfg.BaseEndpoint = &url
	client := apigatewaymanagementapi.NewFromConfig(cfg)

	payload, _ := json.Marshal(data)
	_, err = client.PostToConnection(context.TODO(), &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: connectionId,
		Data:         payload,
	})
	if err != nil {
		return err
	}
	return nil

}
