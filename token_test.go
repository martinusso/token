package token

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestToken(t *testing.T) {
	token := Hash("123", 10)
	got, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		t.Errorf("error %s", err.Error())
	}
	if !strings.HasPrefix(string(got), "$2a$10$") {
		t.Errorf("Invalid prefix")
	}
	if len(got) != 60 {
		t.Errorf("Expected 60 characters', got '%d'", len(got))
	}
}

func TestVerify(t *testing.T) {
	got := Hash("123", 11)
	if !Verify("123", 11, got) {
		t.Errorf("Should return true")
	}
}
