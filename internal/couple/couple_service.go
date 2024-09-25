package couple

import (
	"github.com/cho8833/duary_lambda/internal/auth"
	"github.com/cho8833/duary_lambda/internal/connectcouple"
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
	"math/rand"
	"time"
)

type Service interface {
	CreateCouple(req *CreateCoupleReq, transaction *util.DynamoDBWriteTransaction) (*Couple, util.ApplicationError)
	ConnectCouple(member *auth.LoginMember, coupleCode *string) (*Couple, util.ApplicationError)
}

type ServiceImpl struct {
	repository        Repository
	sessionRepository connectcouple.SessionRepository
}

func NewCoupleService(repository Repository, sessionRepository connectcouple.SessionRepository) *ServiceImpl {
	return &ServiceImpl{repository: repository, sessionRepository: sessionRepository}
}

func (svc *ServiceImpl) CreateCouple(req *CreateCoupleReq, transaction *util.DynamoDBWriteTransaction) (*Couple, util.ApplicationError) {
	isConnected := false
	couple := &Couple{
		IsConnected:    &isConnected,
		RelationDate:   &req.RelationDate,
		OtherCharacter: &req.OtherCharacter,
		Code:           svc.generateCoupleCode(),
	}

	if transaction != nil {
		input, err := svc.repository.GetSaveCoupleTransaction(couple)
		if err != nil {
			log.Printf("failed to get saveCoupleTransaction. error:%s", err.Error())
			return nil, util.DBSaveError{}
		}
		transaction.AddTransaction(input)
		return couple, nil
	} else {
		couple, err := svc.repository.SaveCouple(couple)
		if err != nil {
			log.Printf("failed to save couple\nreq: %+v\nerror:%s", req, err.Error())
			return nil, util.DBSaveError{}
		}
		return couple, nil
	}
}

//func (svc *ServiceImpl) ConnectCouple(member *auth.LoginMember, coupleCode *string) (*Couple, util.ApplicationError) {
//
//}

func (svc *ServiceImpl) generateCoupleCode() *string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano())) // 현재 시간을 시드로 설정
	b := make([]byte, 9)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	result := string(b)
	return &result
}
