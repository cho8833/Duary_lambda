package couple

import "time"

type CreateCoupleReq struct {
	RelationDate   time.Time `json:"relationDate"`
	OtherCharacter string    `json:"otherCharacter"`
}
