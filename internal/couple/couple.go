package couple

import "time"

type Couple struct {
	Id             *string    `json:"id" dynamodbav:"id"`
	IsConnected    *bool      `json:"isConnected" dynamodbav:"isConnected"`
	RelationDate   *time.Time `json:"relationDate" dynamodbav:"relationDate"`
	OtherCharacter *string    `json:"otherCharacter" dynamodbav:"otherCharacter"`
	Code           *string    `json:"code" dynamodb:"code"`
}
