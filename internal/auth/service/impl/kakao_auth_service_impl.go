package impl

import (
	"fmt"
	"github.com/cho8833/duary_lambda/internal/auth/dto"
	"github.com/cho8833/duary_lambda/internal/auth/util"
	"github.com/cho8833/duary_lambda/internal/member/model"
	memberRepository "github.com/cho8833/duary_lambda/internal/member/repository"
	errors "github.com/cho8833/duary_lambda/internal/util"
	"log"
	"os"
)

type KakaoAuthServiceImpl struct {
	memberRepository memberRepository.MemberRepository
	jwtValidator     util.JWTValidator
}

func NewKakaoAuthService(jwtValidator util.JWTValidator,
	memberRepository memberRepository.MemberRepository) *KakaoAuthServiceImpl {
	return &KakaoAuthServiceImpl{jwtValidator: jwtValidator, memberRepository: memberRepository}
}

func (svc *KakaoAuthServiceImpl) SignIn(kakaoToken dto.KakaoOAuthToken) (*dto.SignInRes, error) {

	aud := os.Getenv("aud")
	nonce := os.Getenv("nonce")
	// verify id token
	validateValue := &util.ValidatingValue{
		Url:      "https://kauth.kakao.com/.well-known/jwks.json",
		Aud:      aud,
		Nonce:    nonce,
		Iss:      "https://kauth.kakao.com",
		Provider: "kakao",
	}
	payload, err := svc.jwtValidator.VerifyRSA256(*kakaoToken.IdToken, validateValue)
	if err != nil {
		log.Printf(err.Error())
		return nil, errors.NewBadRequestError(fmt.Errorf("인증정보가 잘못되었습니다"))
	}

	// 카카오 회원 ID 와 카카오 ServiceProvider 로 Member 검색
	member, err := svc.memberRepository.FindBySocialIdAndProvider(payload.SocialId, "kakao")
	if err != nil {
		log.Printf(err.Error())
		return nil, errors.NewDBError(fmt.Errorf("정보를 가져오는데에 실패했습니다"))
	}

	// member 가 존재하는 경우 이미 회원가입된 Member return
	if member != nil {
		result := &dto.SignInRes{
			Member:     member,
			IsRegister: false,
		}
		return result, nil
	} else {
		// member 가 존재하지 않는 경우 Member 생성, 최초 회원가입
		newMember := &model.Member{
			Name:        payload.NickName,
			Birthday:    nil,
			AccessToken: kakaoToken.AccessToken,
			Provider:    "kakao",
			Gender:      nil,
			SocialId:    payload.SocialId,
			FcmToken:    nil,
			Email:       payload.Email,
		}
		err := svc.memberRepository.SaveMember(newMember)
		if err != nil {
			return nil, errors.NewDBError(fmt.Errorf("회원 정보를 저장하는데에 실패했습니다"))
		}
		result := &dto.SignInRes{
			Member:     newMember,
			IsRegister: true,
		}
		return result, nil
	}
}
