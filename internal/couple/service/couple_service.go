package service

import (
	"github.com/cho8833/duary_lambda/internal/couple/model"
	"github.com/cho8833/duary_lambda/internal/couple/repository"
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
)

type CoupleService interface {
	CreateCouple(req *model.CreateCoupleReq) (*model.Couple, util.ApplicationError)
}

type CoupleServiceImpl struct {
	repository repository.CoupleRepository
}

func NewCoupleService(repository repository.CoupleRepository) *CoupleServiceImpl {
	return &CoupleServiceImpl{repository: repository}
}

func (svc *CoupleServiceImpl) CreateCouple(req *model.CreateCoupleReq) (*model.Couple, util.ApplicationError) {
	couple := &model.Couple{
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
