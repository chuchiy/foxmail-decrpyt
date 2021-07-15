package foxmail

import (
	"testing"
)

func TestDecryptPassword(t *testing.T) {
	dec, _ := DecryptPassword("61D05ABC4CFB32CB410C5AFF", true)
	if dec != "smtppass123" {
		t.Fatal("decrpyt error")
	}
	dec, _ = DecryptPassword("279390768847D922F8", true)
	if dec != "password" {
		t.Fatal("decrpyt error")
	}
}
