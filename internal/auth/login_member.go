package auth

import "strings"

type LoginMember struct {
	SocialId string
	Provider string
}

func FromMemberId(memberId *string) *LoginMember {
	s := strings.Split(*memberId, "-")
	return &LoginMember{Provider: s[1], SocialId: s[0]}
}
