package django_models

import (
	"github.com/uptrace/bun"
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
