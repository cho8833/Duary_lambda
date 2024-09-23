package common

import (
	"github.com/cho8833/duary_lambda/internal/couple"
	"github.com/cho8833/duary_lambda/internal/member"
	"github.com/cho8833/duary_lambda/internal/util"
)

type Service interface {
	InitDuaryInfo(request *InitDuaryInfoReq, transaction *util.DynamoDBWriteTransaction) (*InitDuaryInfoRes, util.ApplicationError)
}

type ServiceImpl struct {
	memberSvc member.Service
	coupleSvc couple.Service
}

func NewCommonService(memberSvc member.Service, coupleSvc couple.Service) *ServiceImpl {
	return &ServiceImpl{memberSvc: memberSvc, coupleSvc: coupleSvc}
}

func (svc *ServiceImpl) InitDuaryInfo(request *InitDuaryInfoReq, transaction *util.DynamoDBWriteTransaction) (*InitDuaryInfoRes, util.ApplicationError) {
	// begin transaction
	transaction.BeginTransaction()

	// create new couple
	coupleReq := &couple.CreateCoupleReq{
		RelationDate:   *request.RelationDate,
		OtherCharacter: *request.OtherCharacter,
	}
	newCouple, err := svc.coupleSvc.CreateCouple(coupleReq, transaction)
	if err != nil {
		return nil, err
	}

	// update member
	memberReq := &member.UpdateMemberReq{
		CoupleId:  newCouple.Id,
		Name:      request.Name,
		Birthday:  request.Birthday,
		SocialId:  request.SocialId,
		Character: request.MyCharacter,
		Provider:  request.Provider,
	}
	updatedMember, err := svc.memberSvc.UpdateMember(memberReq, transaction)
	if err != nil {
		return nil, err
	}

	// execute transaction
	_, transactionError := transaction.Execute()
	if transactionError != nil {
		return nil, util.DBError{}
	}

	res := &InitDuaryInfoRes{
		Member: updatedMember,
		Couple: newCouple,
	}
	return res, nil
}
