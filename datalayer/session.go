package datalayer

import (
	"github.com/Flexin1981/gin_django_auth/django_models"
)

type SessionService struct {
}

func (s *SessionService) Get(id string) (*django_models.Session, error) {
	return &django_models.Session{}, nil
}
