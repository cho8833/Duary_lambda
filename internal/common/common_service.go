package common

import (
	"github.com/cho8833/duary_lambda/internal/auth"
	"github.com/cho8833/duary_lambda/internal/connectcouple"
	"github.com/cho8833/duary_lambda/internal/couple"
	"github.com/cho8833/duary_lambda/internal/member"
	"github.com/cho8833/duary_lambda/internal/util"
)

type Service interface {
	InitDuaryInfo(request *InitDuaryInfoReq, transaction *util.DynamoDBWriteTransaction) (*InitDuaryInfoRes, util.ApplicationError)
	ConnectCouple(loginMember *auth.LoginMember, req *ConnectCoupleReq) (*InitDuaryInfoRes, util.ApplicationError)
}

type ServiceImpl struct {
	memberSvc  member.Service
	coupleSvc  couple.Service
	sessionSvc connectcouple.Service
}

func NewCommonService(memberSvc member.Service, coupleSvc couple.Service, sessionSvc connectcouple.Service) *ServiceImpl {
	return &ServiceImpl{memberSvc: memberSvc, coupleSvc: coupleSvc, sessionSvc: sessionSvc}
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

func (svc *ServiceImpl) ConnectCouple(loginMember *auth.LoginMember, req *ConnectCoupleReq, transaction *util.DynamoDBWriteTransaction) (*InitDuaryInfoRes, util.ApplicationError) {
	findMember, svcError := svc.memberSvc.GetMember(loginMember.SocialId, loginMember.Provider)
	if svcError != nil {
		return nil, svcError
	}

	// Find Couple
	findCouple, svcError := svc.coupleSvc.FindByCoupleCode(req.CoupleCode)
	if svcError != nil {
		return nil, svcError
	}

	transaction.BeginTransaction()
	// Update Couple
	connected := true
	findCouple.IsConnected = &connected
	_, svcError = svc.coupleSvc.UpdateCouple(findCouple, transaction)
	if svcError != nil {
		return nil, svcError
	}
	// Update Member
	updateReq := &member.UpdateMemberReq{
		Provider:  loginMember.Provider,
		SocialId:  loginMember.SocialId,
		CoupleId:  findCouple.Id,
		Character: findCouple.OtherCharacter,
	}
	updatedMember, svcError := svc.memberSvc.UpdateMember(updateReq, transaction)
	if svcError != nil {
		return nil, svcError
	}
	_, err := transaction.Execute()
	if err != nil {
		return nil, util.DBUpdateError{}
	}

	// send couple connect message to websocket
	session, svcError := svc.sessionSvc.FindSession(req.CoupleCode)
	if svcError == nil {
		_ = svc.sessionSvc.NotifyCoupleConnected(session, findMember)
	}

	result := &InitDuaryInfoRes{
		Member: updatedMember,
		Couple: findCouple,
	}

	return result, nil
}
