package member

import (
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
)

type Service interface {
	UpdateMember(request *UpdateMemberReq) (*Member, util.ApplicationError)
}

type ServiceImpl struct {
	repo Repository
}

func NewMemberService(repo Repository) *ServiceImpl {
	return &ServiceImpl{repo: repo}
}

func (svc *ServiceImpl) UpdateMember(request *UpdateMemberReq) (*Member, util.ApplicationError) {
	member, err := svc.repo.UpdateMember(request)
	if err != nil {
		log.Printf("failed to update member. req: %+v, error: %s", request, err.Error())
		return nil, util.DBUpdateError{}
	}
	return member, nil
}
