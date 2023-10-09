package django_models

import (
	"github.com/uptrace/bun"
	"math/rand"
	"time"
)

var validDigits = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type (
	Session struct {
		bun.BaseModel `bun:"table:django_session"`

		SessionKey  string    `bun:",pk,session_key"`
		SessionData string    `bun:"session_data"`
		ExpireDate  time.Time `bun:"expire_date"`
	}
)

func NewSession() *Session {
	s := &Session{}
	s.SessionKey = s.CreateKey()
	return s
}

func (s *Session) CreateKey() string {
	b := make([]rune, 32)
	for i := range b {
		b[i] = validDigits[rand.Intn(len(validDigits))]
	}
	return string(b)
}

func (s *Session) EncodeData(data []byte) {

}
