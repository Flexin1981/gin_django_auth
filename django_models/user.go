package django_models

import (
	"github.com/uptrace/bun"
	"time"
)

type (
	AuthUser struct {
		bun.BaseModel `bun:"table:auth_user"`

		Id          int64		`json:"id"`
		Password    string		`json:"password"`
		LastLogin   time.Time	`json:"last_login"`
		Username    string		`json:"username"`
		FirstName   string		`json:"first_name"`
		LastName    string		`json:"last_name"`
		Email       string		`json:"email"`
		IsStaff     bool		`json:"is_staff"`
		IsActive    bool		`json:"is_active"`
		DateJoined  time.Time	`json:"date_joined"`
	}
)
