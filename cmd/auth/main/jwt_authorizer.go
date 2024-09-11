package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/duary_lambda/internal/auth/jwtutil"
	"log"
	"os"
)

func jwtAuthorizer(ctx context.Context, request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayV2CustomAuthorizerIAMPolicyResponse, error) {
	if request.Type != "Bearer" {
		return denyResponse(), nil
	}
	key := os.Getenv("secretKey")
	jwtUtil := jwtutil.JWTUtilImpl{}
	id, err := jwtUtil.ValidateApplicationJWT(request.AuthorizationToken, key)
	if err != nil {
		log.Printf("Authorization: fail to authorize. token: %s, error: %s", request.AuthorizationToken, err.Error())
		return denyResponse(), nil
	}
	return events.APIGatewayV2CustomAuthorizerIAMPolicyResponse{
		PrincipalID: *id,
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
		},
	}, nil
}

func denyResponse() events.APIGatewayV2CustomAuthorizerIAMPolicyResponse {
	return events.APIGatewayV2CustomAuthorizerIAMPolicyResponse{
		PrincipalID: "",
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   "Deny",
					Resource: []string{"*"},
				},
			},
		},
	}
}

func main() {
	lambda.Start(jwtAuthorizer)
}
