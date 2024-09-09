package dto

import (
	"github.com/cho8833/duary_lambda/internal/auth/jwt_util"
	"github.com/cho8833/duary_lambda/internal/member/model"
)

type SignInRes struct {
	IsRegister bool
	Member     *model.Member
	Token      *jwt_util.ApplicationJWT
}
