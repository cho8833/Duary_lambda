package model

type OAuthToken struct {
	accessToken *string
	provider    *string
	memberId    *int64
	socialId    *int64
}
