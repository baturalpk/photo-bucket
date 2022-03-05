package entclient

import (
	"testing"
)

func TestInitConnection(t *testing.T) {
	if err := InitConnection(); err != nil {
		t.Error(err)
		t.Fail()
	}
}
