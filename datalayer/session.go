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

type SessionService struct {
}

func (s *SessionService) Get(id string) (*django_models.Session, error) {
	var djangoSession django_models.Session
	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv(DatabaseConnectionEnvironmentVariable)))), pgdialect.New())
	if err := db.NewSelect().Model(&djangoSession).Where(BunQueryString, bun.Ident("id"), id).Scan(context.Background()); err != nil {
		return &djangoSession, err
	}
	return &djangoSession, nil
}

func (s *SessionService) Create(user *django_models.AuthUser) (sessionId string, err error) {
	var djangoSession django_models.Session
	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv(DatabaseConnectionEnvironmentVariable)))), pgdialect.New())
	if err := db.NewSelect().Model(&djangoSession).Where(BunQueryString, bun.Ident("id"), id).Scan(context.Background()); err != nil {
		return djangoSession.SessionKey, err
	}
	return djangoSession.SessionKey, nil
}
