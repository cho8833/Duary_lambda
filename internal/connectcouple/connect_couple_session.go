package connectcouple

type Session struct {
	CoupleCode   *string `dynamodbav:"coupleCode"`
	MemberId     *string `dynamodbav:"memberId"`
	CoupleId     *string `dynamodbav:"coupleId"`
	ConnectionId *string `dynamodbav:"connectionId"`
}
