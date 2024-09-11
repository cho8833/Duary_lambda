package service

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cho8833/duary_lambda/internal/auth/dto"
	"github.com/cho8833/duary_lambda/internal/auth/jwtutil"
	"github.com/cho8833/duary_lambda/internal/member/model"
	"github.com/cho8833/duary_lambda/internal/member/repository"
	"github.com/cho8833/duary_lambda/internal/util"
	"log"
	"os"
)

type KakaoAuthService interface {
	SignIn(token *dto.KakaoOAuthToken) (dto.SignInRes, util.ApplicationError)
}

type KakaoAuthServiceImpl struct {
	memberRepository repository.MemberRepository
	jwtValidator     jwtutil.JWTValidator
	jwtUtil          jwtutil.JWTUtil
}

func NewKakaoAuthService(jwtValidator jwtutil.JWTValidator,
	jwtUtil jwtutil.JWTUtil,
	memberRepository repository.MemberRepository) *KakaoAuthServiceImpl {
	return &KakaoAuthServiceImpl{jwtValidator: jwtValidator, jwtUtil: jwtUtil, memberRepository: memberRepository}
}

func (svc *KakaoAuthServiceImpl) SignIn(kakaoToken *dto.KakaoOAuthToken) (*dto.SignInRes, util.ApplicationError) {

	aud := os.Getenv("aud")
	nonce := os.Getenv("nonce")
	// verify id token
	validateValue := &jwtutil.ValidatingValue{
		Url:      "https://kauth.kakao.com/.well-known/jwks.json",
		Aud:      aud,
		Nonce:    nonce,
		Iss:      "https://kauth.kakao.com",
		Provider: "kakao",
	}
	payload, err := svc.jwtValidator.VerifyRSA256(*kakaoToken.IdToken, validateValue)
	if err != nil {
		log.Printf("failed to verify token. idToken: %s, error: %s", kakaoToken.IdToken, err.Error())
		return nil, util.BadRequestError{}
	}

	// 카카오 회원 ID 와 카카오 ServiceProvider 로 Member 검색
	// Member 가 없을 경우 ResourceNotFoundException 발생, 해당 Exception 은 오류가 아님
	member, err := svc.memberRepository.FindBySocialIdAndProvider(payload.SocialId, "kakao")
	if temp := new(types.ResourceNotFoundException); !errors.As(err, &temp) && err != nil {
		id := fmt.Sprintf("%dkakao", payload.SocialId)
		log.Printf("failed to find member\nid:%s\nerror:%s", id, err.Error())
		return nil, util.DBReadError{}
	}

	// generate application token
	memberId := fmt.Sprintf("%d%s", member.SocialId, member.Provider)
	key := os.Getenv("secretKey")
	newToken := svc.jwtUtil.NewToken(memberId, key)

	// member 가 존재하는 경우 DB 필드를 업데이트하고 이미 회원가입된 Member return
	if member != nil {
		member.AccessToken = kakaoToken.AccessToken
		_, err := svc.memberRepository.SaveMember(member)
		if err != nil {
			log.Printf("failed to save member\nmember: %+v\nerror: %s", member, err.Error())
			return nil, util.DBSaveError{}
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
		_, err := svc.memberRepository.SaveMember(newMember)
		if err != nil {
			log.Printf("failed to save member\nnew member: %+v\nerror: %s", newMember, err.Error())
			return nil, util.DBSaveError{}
		}
		result := &dto.SignInRes{
			Member:     newMember,
			IsRegister: true,
			Token:      newToken,
		}
		return result, nil
	}
}
