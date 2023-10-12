package datalayer

import (
	"errors"
	"github.com/Flexin1981/gin_django_auth/django_models"
)

type MockSessionService struct {
}

func (s *MockSessionService) Create(user *django_models.AuthUser) (*django_models.Session, error) {
	return &django_models.Session{}, nil
}

func (s *MockSessionService) Get(id string) (*django_models.Session, error) {
	if id == "error" {
		return nil, errors.New("unknown id")
	}
	return &django_models.Session{}, nil
}
