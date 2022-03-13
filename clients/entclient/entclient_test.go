package entclient

import (
	"os"
	"testing"
)

func TestInitConnection(t *testing.T) {
	os.Setenv("env", "test")
	if err := InitConnection(); err != nil {
		t.Error(err)
		t.Fail()
	}
}
