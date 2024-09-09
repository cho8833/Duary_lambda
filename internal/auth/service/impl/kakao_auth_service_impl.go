package impl

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cho8833/duary_lambda/internal/auth/dto"
	"github.com/cho8833/duary_lambda/internal/auth/jwt_util"
	"github.com/cho8833/duary_lambda/internal/member/model"
	memberRepository "github.com/cho8833/duary_lambda/internal/member/repository"
	appError "github.com/cho8833/duary_lambda/internal/util"
	"log"
	"os"
)

type KakaoAuthServiceImpl struct {
	memberRepository memberRepository.MemberRepository
	jwtValidator     jwt_util.JWTValidator
	jwtUtil          jwt_util.JWTUtil
}

func NewKakaoAuthService(jwtValidator jwt_util.JWTValidator,
	jwtUtil jwt_util.JWTUtil,
	memberRepository memberRepository.MemberRepository) *KakaoAuthServiceImpl {
	return &KakaoAuthServiceImpl{jwtValidator: jwtValidator, jwtUtil: jwtUtil, memberRepository: memberRepository}
}

func (svc *KakaoAuthServiceImpl) SignIn(kakaoToken *dto.KakaoOAuthToken) (*dto.SignInRes, appError.ApplicationError) {

	aud := os.Getenv("aud")
	nonce := os.Getenv("nonce")
	// verify id token
	validateValue := &jwt_util.ValidatingValue{
		Url:      "https://kauth.kakao.com/.well-known/jwks.json",
		Aud:      aud,
		Nonce:    nonce,
		Iss:      "https://kauth.kakao.com",
		Provider: "kakao",
	}
	payload, err := svc.jwtValidator.VerifyRSA256(*kakaoToken.IdToken, validateValue)
	if err != nil {
		log.Printf(err.Error())
		return nil, appError.NewBadRequestError(fmt.Errorf("인증정보가 잘못되었습니다"))
	}

	// 카카오 회원 ID 와 카카오 ServiceProvider 로 Member 검색
	// Member 가 없을 경우 ResourceNotFoundException 발생, 해당 Exception 은 오류가 아님
	member, err := svc.memberRepository.FindBySocialIdAndProvider(payload.SocialId, "kakao")
	if temp := new(types.ResourceNotFoundException); !errors.As(err, &temp) && err != nil {
		log.Printf(err.Error())
		return nil, appError.NewDBError(fmt.Errorf("정보를 가져오는데에 실패했습니다"))
	}

	// generate application token
	memberId := fmt.Sprintf("%d%s", member.SocialId, member.Provider)
	newToken := svc.jwtUtil.NewToken(memberId, *member.Name)

	// member 가 존재하는 경우 DB 필드를 업데이트하고 이미 회원가입된 Member return
	if member != nil {
		member.AccessToken = kakaoToken.AccessToken
		err := svc.memberRepository.SaveMember(member)
		if err != nil {
			return nil, appError.NewDBError(fmt.Errorf("회원 정보를 저장하는데에 실패했습니다"))
		}
		result := &dto.SignInRes{
			Member:     member,
			IsRegister: false,
			Token:      newToken,
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
			return nil, appError.NewDBError(fmt.Errorf("회원 정보를 저장하는데에 실패했습니다"))
		}
		result := &dto.SignInRes{
			Member:     newMember,
			IsRegister: true,
			Token:      newToken,
		}
		return result, nil
	}
}
