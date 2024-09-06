package service

import (
	"github.com/cho8833/Duary/internal/util"
)

type OIDCService interface {
	GetPublicKey(url string, provider string, kid string) (*util.ServerResponse, error)
}
