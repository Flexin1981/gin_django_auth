package datalayer

import (
	"gin_django_auth/django_models"
	"os"
)

type SessionServiceInterface interface {
	Get(id string) (*django_models.Session, error)
}

func NewSessionService() SessionServiceInterface {
	if os.Getenv("GINDJANGOAUTHTEST") == "true" {
		return &MockSessionService{}
	}
	return &SessionService{}
}
