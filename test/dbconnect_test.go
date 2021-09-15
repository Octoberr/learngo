package test

import (
	"learngo/dbhelper"
	"testing"
)

func Test_DBconnect(t *testing.T) {
	dbhelper.InitDB()
	t.Log("test db connect")
}
