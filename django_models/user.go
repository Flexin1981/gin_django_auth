package django_models

import (
	"crypto/sha1"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/pbkdf2"
	"strconv"
	"strings"
	"time"
)

type (
	AuthUser struct {
		bun.BaseModel `bun:"table:auth_user"`

		id          int64
		Password    string
		LastLogin   time.Time
		IsSuperUser bool
		Username    string
		Firstname   string
		LastName    string
		Email       string
		IsStaff     bool
		IsActive    bool
		DateJoined  time.Time
	}
)

func (a *AuthUser) Encode(password, salt string, iterations int) string {
	return string(pbkdf2.Key([]byte(password), []byte(salt), iterations, 32, sha1.New))
}

func (a *AuthUser) Verify(password string) bool {
	splitEncoded := strings.Split(a.Password, "$")
	iterations, _ := strconv.ParseInt(splitEncoded[1], 10, 64)
	return a.Encode(password, splitEncoded[2], int(iterations)) == a.Password
}
