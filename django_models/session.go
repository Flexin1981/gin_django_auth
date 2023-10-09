package django_models

import (
	"github.com/uptrace/bun"
	"time"
)

type (
	Session struct {
		bun.BaseModel `bun:"table:django_session"`

		SessionKey  string    `bun:",pk,session_key"`
		SessionData string    `bun:"session_data"`
		ExpireDate  time.Time `bun:"expire_date"`
	}
)

func (s *Session) CreateKey() {

}

func (s *Session) EncodeData(data []byte) {

}
