package datalayer

import (
	"context"
	"github.com/Flexin1981/gin_django_auth/django_models"
	"github.com/uptrace/bun"
)

type AuthUserService struct {
}

func (s *AuthUserService) Get(id string) (*django_models.AuthUser, error) {
	var djangoUser django_models.AuthUser
	db := GetDatabaseConnection()
	if err := db.NewSelect().Model(&djangoUser).Where(BunQueryString, bun.Ident("id"), id).Scan(context.Background()); err != nil {
		return &djangoUser, err
	}
	return &djangoUser, nil
}

func (s *AuthUserService) GetByUsername(username string) (*django_models.AuthUser, error) {
	var djangoUser django_models.AuthUser
	db := GetDatabaseConnection()
	if err := db.NewSelect().Model(&djangoUser).Where(BunQueryString, bun.Ident("username"), username).Scan(context.Background()); err != nil {
		return &djangoUser, err
	}
	return &djangoUser, nil
}
