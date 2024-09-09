package service

import (
	"github.com/cho8833/duary_lambda/internal/util"
)

type OIDCService interface {
	GetPublicKey(url string, provider string, kid string) (*util.ServerResponse, error)
}
