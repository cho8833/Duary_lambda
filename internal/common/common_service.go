package common

import (
	"github.com/cho8833/duary_lambda/internal/couple"
	"github.com/cho8833/duary_lambda/internal/member"
	"github.com/cho8833/duary_lambda/internal/util"
)

type Service interface {
	InitDuaryInfo(request InitDuaryInfoReq) (*InitDuaryInfoRes, util.ApplicationError)
}

type ServiceImpl struct {
	memberSvc member.Service
	coupleSvc couple.Service
}

func NewCommonService(memberSvc member.Service, coupleSvc couple.Service) *ServiceImpl {
	return &ServiceImpl{memberSvc: memberSvc, coupleSvc: coupleSvc}
}

func (svc *ServiceImpl) InitDuaryInfo(request InitDuaryInfoReq) (*InitDuaryInfoRes, util.ApplicationError) {
	// create new couple
	coupleReq := &couple.CreateCoupleReq{
		RelationDate:   request.RelationDate,
		OtherCharacter: request.OtherCharacter,
	}
	newCouple, err := svc.coupleSvc.CreateCouple(coupleReq)
	if err != nil {
		return nil, err
	}

	memberReq := &member.UpdateMemberReq{
		CoupleId: &newCouple.Id,
		Name:     &request.Name,
		Birthday: &request.Birthday,
	}
	updatedMember, err := svc.memberSvc.UpdateMember(memberReq)
	if err != nil {
		return nil, err
	}

	res := &InitDuaryInfoRes{
		Member: updatedMember,
		Couple: newCouple,
	}
	return res, nil
}
