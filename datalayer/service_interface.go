package datalayer

import (
	"github.com/Flexin1981/gin_django_auth/django_models"
	"os"
)

type SessionServiceInterface interface {
	Get(id string) (*django_models.Session, error)
	Create(user *django_models.AuthUser) (string, error)
}

type AuthUserServiceInterface interface {
	Get(id string) (*django_models.AuthUser, error)
	GetByUsername(username string) (*django_models.AuthUser, error)
}

func NewSessionService() SessionServiceInterface {
	if os.Getenv("GINDJANGOAUTHTEST") == "true" {
		//ToDo: fix the mock testing once this is up and running
		return nil
	}
	return &SessionService{}
}

func NewAuthUserService() AuthUserServiceInterface {
	if os.Getenv("GINDJANGOAUTHTEST") == "true" {
		//ToDo: fix the mock testing once this is up and running
		return nil
	}
	return &AuthUserService{}
}
