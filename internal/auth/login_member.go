package auth

import (
	"github.com/aws/aws-lambda-go/events"
	"strconv"
	"strings"
)

type LoginMember struct {
	SocialId int64
	Provider string
}

func FromMemberId(memberId *string) *LoginMember {
	s := strings.Split(*memberId, "-")
	socialId, _ := strconv.ParseInt(s[0], 10, 64)
	return &LoginMember{Provider: s[1], SocialId: socialId}
}

func FromRequestContext(req events.APIGatewayProxyRequest) *LoginMember {
	lambdaMap := req.RequestContext.Authorizer["lambda"].(map[string]interface{})
	provider := lambdaMap["provider"].(string)
	socialId, _ := strconv.ParseInt(lambdaMap["socialId"].(string), 10, 64)
	return &LoginMember{Provider: provider, SocialId: socialId}
}
