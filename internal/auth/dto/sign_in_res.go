package dto

import "github.com/cho8833/Duary/internal/member/model"

type SignInRes struct {
	IsRegister bool
	Member     *model.Member
}
