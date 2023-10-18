package django_models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func containsOnly(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestSessionCookieIsCorrectLength(t *testing.T) {
	session := Session{}
	assert.Len(t, session.CreateKey(), 32)

}

func TestSessionCookieContainsCorrectDigits(t *testing.T) {
	session := Session{}
	key := session.CreateKey()
	for _, letter := range []rune(key) {
		if !containsOnly(validDigits, letter) {
			t.Error(letter)
		}
	}
}

func TestSessionCookieEncodesCorrectly(t *testing.T) {
	session := Session{}
	
	t.Setenv("GINDJANGOAUTHSIGNINGKEY", "django-insecure-2z74-f1&^1xqaiw4!9^32^c*(9zr-zs1y5w2j9dlftb4@rz_f5")
	
	signedObj := session.SignObject([]byte(`{"_auth_user_hash":"39308b9542b9305fc038d28a51088905e14246a1","_auth_user_backend":"x.alternate_auth.Backend","_auth_user_id":"52135"}`))

	if strings.Split(signedObj, `:`)[0] !=  `.eJxUy8EKwjAMgOF3yVlGmjSS7OiLlHSLVJQdtg4E8d1F8aDn__sfUHzvrexbrKX51mAENkatJpmqMcp5QtaZ1CWhqqFEypSPnuDwO1efrrHMMMJ98FuPdfEenz6cvunPX95UKLHA8xUAAP__-1AqZg` {
		t.Error("data not encoded correctly")
	}
}
