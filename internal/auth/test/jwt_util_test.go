package test

import (
	"github.com/cho8833/duary_lambda/internal/auth/jwt_util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_NewToken(t *testing.T) {
	os.Setenv("secretKey", "4fM090neAzdhL/RRduD0M8IEBgY7N+uR/7pJmO92wWo=")
	jwtUtil := jwt_util.JWTUtilImpl{}

	applicationJwt := jwtUtil.NewToken("1kakao", "조현빈")
	id, err := jwtUtil.ValidateApplicationJWT(applicationJwt.AccessToken)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, "1kakao", *id)
	id, err = jwtUtil.ValidateApplicationJWT(applicationJwt.RefreshToken)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, "1kakao", *id)
}

func Test_ValidateApplicationJWT(t *testing.T) {
	os.Setenv("secretKey", "4fM090neAzdhL/RRduD0M8IEBgY7N+uR/7pJmO92wWo=")
	jwtUtil := jwt_util.JWTUtilImpl{}

	applicationJwt := jwtUtil.NewToken("1kakao", "조현빈")

	id, err := jwtUtil.ValidateApplicationJWT(applicationJwt.AccessToken)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, "1kakao", *id)
}
