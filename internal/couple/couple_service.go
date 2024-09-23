package couple

import (
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
)

type Service interface {
	CreateCouple(req *CreateCoupleReq, transaction *util.DynamoDBWriteTransaction) (*Couple, util.ApplicationError)
}

type ServiceImpl struct {
	repository Repository
}

func NewCoupleService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

func (svc *ServiceImpl) CreateCouple(req *CreateCoupleReq, transaction *util.DynamoDBWriteTransaction) (*Couple, util.ApplicationError) {
	isConnected := false
	couple := &Couple{
		IsConnected:    &isConnected,
		RelationDate:   &req.RelationDate,
		OtherCharacter: &req.OtherCharacter,
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
