package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cho8833/duary_lambda/internal/auth/jwtutil"
	"log"
	"os"
	"strings"
)

/*
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -tags lambda.norpc -o bootstrap cmd/auth/main/jwt_authorizer.go && chmod 755 bootstrap && zip  build/package/jwt_authorizer.zip bootstrap && rm bootstrap
*/
func jwtAuthorizer(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayV2CustomAuthorizerIAMPolicyResponse, error) {
	key := os.Getenv("secretKey")
	jwtUtil := jwtutil.Impl{}
	header := request.Headers["authorization"]
	log.Printf("got authorization : %s", header)

	check := strings.Split(header, " ")
	token := check[1]
	id, err := jwtUtil.ValidateApplicationJWT(token, key)
	if err != nil {
		log.Printf("Authorization: fail to authorize. token: %s, error: %s", token, err.Error())
		return denyResponse(), nil
	}
	s := strings.Split(*id, "-")

	log.Printf("authorized %s", *id)
	return events.APIGatewayV2CustomAuthorizerIAMPolicyResponse{
		PrincipalID: *id,
		Context: map[string]interface{}{
			"socialId": s[0],
			"provider": s[1],
		},
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action: []string{"execute-api:Invoke"},
					Effect: "Allow",
					Resource: []string{
						"arn:aws:execute-api:ap-northeast-2:922001515124:i3lm91q9v5/*",
					},
				},
			},
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
