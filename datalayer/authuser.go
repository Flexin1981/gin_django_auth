package datalayer

import (
	"context"
	"database/sql"
	"github.com/Flexin1981/gin_django_auth/django_models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

type AuthUserService struct {
}

func (s *AuthUserService) Get(id string) (*django_models.AuthUser, error) {
	var djangoUser django_models.AuthUser
	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv(DatabaseConnectionEnvironmentVariable)))), pgdialect.New())
	if err := db.NewSelect().Model(&djangoUser).Where(BunQueryString, bun.Ident("id"), id).Scan(context.Background()); err != nil {
		return &djangoUser, err
	}
	return &djangoUser, nil
}

func (s *AuthUserService) GetByUsername(username string) (*django_models.AuthUser, error) {
	var djangoUser django_models.AuthUser
	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv(DatabaseConnectionEnvironmentVariable)))), pgdialect.New())
	if err := db.NewSelect().Model(&djangoUser).Where(BunQueryString, bun.Ident("username"), username).Scan(context.Background()); err != nil {
		return &djangoUser, err
	}
	return &djangoUser, nil
}
