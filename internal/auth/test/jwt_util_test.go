package test

import (
	"github.com/cho8833/duary_lambda/internal/auth/jwtutil"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_NewToken(t *testing.T) {
	key := "4fM090neAzdhL/+uR/7pJmO92wWo="
	os.Setenv("secretKey", key)
	jwtUtil := jwtutil.JWTUtilImpl{}

	applicationJwt := jwtUtil.NewToken("1kakao", "조현빈")
	id, err := jwtUtil.ValidateApplicationJWT(applicationJwt.AccessToken, key)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, "1kakao", *id)
	id, err = jwtUtil.ValidateApplicationJWT(applicationJwt.RefreshToken, key)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, "1kakao", *id)
}

func Test_ValidateApplicationJWT(t *testing.T) {
	key := "4fM090neAzdhL/+uR/7pJmO92wWo="
	os.Setenv("secretKey", key)
	jwtUtil := jwtutil.JWTUtilImpl{}

	applicationJwt := jwtUtil.NewToken("1kakao", "조현빈")

	id, err := jwtUtil.ValidateApplicationJWT(applicationJwt.AccessToken, key)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, "1kakao", *id)
}
