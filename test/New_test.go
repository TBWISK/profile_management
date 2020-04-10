package test

import (
	"fmt"
	"tbwisk/dao"
	"testing"
)

func Test_uuid(t *testing.T) {
	fmt.Println(dao.NewUUID())
}
