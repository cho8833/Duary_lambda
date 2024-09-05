package model

type OAuthToken struct {
	AccessToken *string `dynamodbav:"accessToken"`
	Provider    *string `dynamodbav:"provider"`
	MemberId    *int64  `dynamodbav:"memberId"`
	SocialId    *int64  `dynamodbav:"socialId"`
}
