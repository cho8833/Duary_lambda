package connectcouple

type SessionReq struct {
	CoupleCode   *string `json:"coupleCode"`
	MemberId     *string `json:"memberId"`
	CoupleId     *string `json:"coupleId"`
	ConnectionId *string
}
