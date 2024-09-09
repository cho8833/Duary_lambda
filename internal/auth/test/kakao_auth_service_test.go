package test

//
//import (
//	"errors"
//	"fmt"
//	"github.com/cho8833/duary_lambda/internal/auth/dto"
//	"github.com/cho8833/duary_lambda/internal/auth/service/impl"
//	jwtUtil "github.com/cho8833/duary_lambda/internal/auth/util"
//	customErrors "github.com/cho8833/duary_lambda/internal/util"
//	"github.com/cho8833/duary_lambda/mocks"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func Test_SignIn(t *testing.T) {
//	mockMemberRepo := mocks.NewMemberRepository(t)
//
//	dummyValidatingValue := jwtUtil.ValidatingValue{
//		Iss:      "dummyIss",
//		Aud:      "dummyAud",
//		Nonce:    "dummyNonce",
//		Url:      "dummyUrl",
//		Provider: "dummyProvider",
//	}
//	t.Run("idToken Validating 에 실패할 경우 BadRequestError 반환", func(t *testing.T) {
//		// given
//		mockValidator.EXPECT().VerifyRSA256(gomock.Any(), dummyValidatingValue).Return(
//			nil, fmt.Errorf("dummy error"))
//
//		// when
//		target := impl.NewKakaoAuthService(mockValidator, mockMemberRepo)
//		res, err := target.SignIn(dto.KakaoOAuthToken{})
//
//		// then
//		assert.Nil(t, res)
//		var v customErrors.BadRequestError
//		ok := errors.As(err, &v)
//		assert.Equal(t, ok, true)
//
//	})
//}
