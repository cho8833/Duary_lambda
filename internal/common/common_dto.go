package common

import (
	"github.com/cho8833/duary_lambda/internal/couple"
	"github.com/cho8833/duary_lambda/internal/member"
	"github.com/cho8833/duary_lambda/internal/util"
	"time"
)

type InitDuaryInfoReq struct {
	Name           *string    `json:"name"`
	Birthday       *time.Time `json:"birthday"`
	RelationDate   *time.Time `json:"relationDate"`
	OtherCharacter *string    `json:"otherCharacter"`
	MyCharacter    *string    `json:"myCharacter"`
	Provider       string
	SocialId       int64
}

func (req InitDuaryInfoReq) Validate() util.ApplicationError {
	if req.OtherCharacter == req.MyCharacter {
		return util.NewCustomApplicationError("동일한 캐릭터를 사용할 수 없습니다")
	}
	return nil
}

type InitDuaryInfoRes struct {
	Member *member.Member `json:"member"`
	Couple *couple.Couple `json:"couple"`
}
