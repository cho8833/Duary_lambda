package connectcouple

import (
	"github.com/cho8833/duary_lambda/internal/member"
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
)

type Service interface {
	CreateSession(req *SessionReq) (*Session, util.ApplicationError)
	DeleteSession(requestId *string) util.ApplicationError
	FindSession(coupleCode *string) (*Session, util.ApplicationError)
	NotifyCoupleConnected(session *Session, connectedMember *member.Member) util.ApplicationError
}

type ServiceImpl struct {
	repository           SessionRepository
	apiGatewayRepository ApiGatewayRepository
}

func NewConnectCoupleService(repository SessionRepository) *ServiceImpl {
	return &ServiceImpl{repository: repository, apiGatewayRepository: &ApiGatewayRepositoryImpl{}}
}

func (svc *ServiceImpl) CreateSession(req *SessionReq) (*Session, util.ApplicationError) {
	session := &Session{
		CoupleCode:   req.CoupleCode,
		MemberId:     req.MemberId,
		CoupleId:     req.CoupleId,
		ConnectionId: req.ConnectionId,
	}
	result, err := svc.repository.SaveSession(session)
	if err != nil {
		log.Printf("failed to save session. error: %s", err.Error())
		return nil, util.DBSaveError{}
	}
	return result, nil
}

func (svc *ServiceImpl) FindSession(coupleCode *string) (*Session, util.ApplicationError) {
	sessions, err := svc.repository.FindByCoupleCode(coupleCode)
	if err != nil {
		return nil, util.DBReadError{}
	}
	if len(sessions) != 1 {
		log.Printf("Session with coupleCode: %s invalid. found %d", *coupleCode, len(sessions))
		return nil, util.InternalServerError{}
	}
	return &sessions[0], nil
}

func (svc *ServiceImpl) DeleteSession(requestId *string) util.ApplicationError {
	err := svc.repository.DeleteByConnectionId(requestId)
	if err != nil {
		log.Printf("failed to delete session. error: %s", err.Error())
		return util.DBDeleteError{}
	}
	return nil
}

func (svc *ServiceImpl) NotifyCoupleConnected(session *Session, connectedMember *member.Member) util.ApplicationError {
	url := "https://2t0inm29d4.execute-api.ap-northeast-2.amazonaws.com/dev"
	err := svc.apiGatewayRepository.PostToConnect(url, session.ConnectionId, connectedMember)
	if err != nil {
		log.Printf("failed to post to connect. session: %+v\nerror: %s", session, err.Error())
		return util.InternalServerError{}
	}
	return nil

}
