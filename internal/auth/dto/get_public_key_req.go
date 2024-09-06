package dto

type GetPublicKeyReq struct {
	Url      string `json:"url"`
	Provider string `json:"provider"`
	Kid      string `json:"kid"`
}
