package couple

import "time"

type CreateCoupleReq struct {
	RelationDate   time.Time `json:"relationDate"`
	OtherCharacter string    `json:"otherCharacter"`
}

type UpdateCoupleReq struct {
	IsConnected *bool   `dynamodbav:"isConnected"`
	Id          *string `dynamodbav:"id"`
}
