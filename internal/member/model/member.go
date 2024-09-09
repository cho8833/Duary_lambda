package model

import "time"

type Member struct {
	Name        *string    `json:"name" dynamodbav:"name""`
	Birthday    *time.Time `json:"birthday" dynamodbav:"birthday"`
	Gender      *string    `json:"gender" dynamodbav:"gender"` // man, woman, other
	FcmToken    *string    `json:"fcmToken" dynamodbav:"fcmToken"`
	AccessToken *string    `dynamodbav:"accessToken"`
	Provider    string     `dynamodbav:"provider"`
	SocialId    int64      `json:"socialId" dynamodbav:"socialId"`
	Email       *string    `json:"email" dynamodbav:"email"`
}
