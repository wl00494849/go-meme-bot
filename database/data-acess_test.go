package database

import "testing"

func Test_DBInit(t *testing.T) {
	if err := DBInit(); err != nil {
		t.Log(err)
	}
}
