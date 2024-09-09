package repository

import "github.com/cho8833/duary_lambda/internal/member/model"

type MemberRepository interface {
	FindBySocialIdAndProvider(socialId int64, provider string) (*model.Member, error)
	SaveMember(member *model.Member) error
}
