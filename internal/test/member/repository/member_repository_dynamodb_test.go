package repository

import (
	"fmt"
	"github.com/cho8833/Duary/internal/member/model"
	"github.com/cho8833/Duary/internal/member/repository/impl"
	"github.com/cho8833/Duary/internal/test/util"
	"testing"
	"time"
)

func Test_FindById(t *testing.T) {
	client := util.CreateLocalDynamoDBClient()
	repository := impl.NewMemberRepository(client)

	member, err := repository.FindBySocialIdAndProvider(1, "")
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Printf("%+v", member)
}

func Test_saveMember(t *testing.T) {
	client := util.CreateLocalDynamoDBClient()
	repository := impl.NewMemberRepository(client)
	name := "test"
	now := time.Now()
	gender := "man"
	dummyMember := &model.Member{
		SocialId: 1,
		Name:     &name,
		Birthday: &now,
		Gender:   &gender,
		Provider: "kakao",
		FcmToken: nil,
	}

	err := repository.SaveMember(dummyMember)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func Test_findBySocialIdAndProvider(t *testing.T) {
	client := util.CreateLocalDynamoDBClient()
	repository := impl.NewMemberRepository(client)

	member, err := repository.FindBySocialIdAndProvider(1, "kakao")
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Printf("%+v", member)
}
