package model

import "time"

type Couple struct {
	Id             string    `json:"id"`
	IsConnected    bool      `json:"isConnected"`
	RelationDate   time.Time `json:"relationDate"`
	OtherCharacter string    `json:"otherCharacter"`
}

type CreateCoupleReq struct {
	RelationDate   time.Time `json:"relationDate"`
	OtherCharacter string    `json:"otherCharacter"`
}
