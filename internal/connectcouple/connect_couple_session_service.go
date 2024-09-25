package connectcouple

import (
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
)

type Service interface {
	CreateSession(req *SessionReq) (*Session, util.ApplicationError)
	DeleteSession(requestId *string) util.ApplicationError
}

type ServiceImpl struct {
	repository SessionRepository
}

func NewWebsocketHandler(repository SessionRepository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

func (handler *ServiceImpl) CreateSession(req *SessionReq) (*Session, util.ApplicationError) {
	session := &Session{
		CoupleCode:   req.CoupleCode,
		MemberId:     req.MemberId,
		CoupleId:     req.CoupleId,
		ConnectionId: req.ConnectionId,
	}
	result, err := handler.repository.SaveSession(session)
	if err != nil {
		log.Printf("failed to save session. error: %s", err.Error())
		return nil, util.DBSaveError{}
	}
	return result, nil
}

func (handler *ServiceImpl) DeleteSession(requestId *string) util.ApplicationError {
	err := handler.repository.DeleteByConnectionId(requestId)
	if err != nil {
		log.Printf("failed to delete session. error: %s", err.Error())
		return util.DBDeleteError{}
	}
	return nil
}
