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

		Id          int64		`json:"id"`
		Password    string		`json:"password"`
		LastLogin   time.Time	`json:"last_login"`
		IsSuperUser bool		`json:"is_super_user"`
		Username    string		`json:"username"`
		Firstname   string		`json:"firstname"`
		LastName    string		`json:"last_name"`
		Email       string		`json:"email"`
		IsStaff     bool		`json:"is_staff"`
		IsActive    bool		`json:"is_active"`
		DateJoined  time.Time	`json:"date_joined"`
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
