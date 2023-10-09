package django_models

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
