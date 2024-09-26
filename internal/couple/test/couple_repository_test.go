package test

import (
	"fmt"
	"github.com/cho8833/duary_lambda/internal/couple"
	"github.com/cho8833/duary_lambda/internal/test/util"
	"testing"
)

func Test_FindByCoupleCode(t *testing.T) {
	dynamodbClient := util.CreateLocalDynamoDBClient()

	repository := couple.NewCoupleRepository(dynamodbClient)

	coupleCode := "72itgwj0t"
	couples, err := repository.FindByCoupleCode(&coupleCode)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Printf("%+v", couples)
}
