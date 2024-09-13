package couple

import (
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
)

type Service interface {
	CreateCouple(req *CreateCoupleReq) (*Couple, util.ApplicationError)
}

type ServiceImpl struct {
	repository Repository
}

func NewCoupleService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

func (svc *ServiceImpl) CreateCouple(req *CreateCoupleReq) (*Couple, util.ApplicationError) {
	couple := &Couple{
		IsConnected:    false,
		RelationDate:   req.RelationDate,
		OtherCharacter: req.OtherCharacter,
	}

	couple, err := svc.repository.SaveCouple(couple)
	if err != nil {
		log.Printf("failed to save couple\nreq: %+v\nerror:%s", req, err.Error())
		return nil, util.DBSaveError{}
	}

	return couple, nil
}
