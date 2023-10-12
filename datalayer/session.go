package datalayer

import (
	"fmt"
	"context"
	"database/sql"
	"encoding/json"
	"os"

	"github.com/Flexin1981/gin_django_auth/django_models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type (

	SessionData struct {
		AuthUserHash 	string `json:"_auth_user_hash"`
		AuthUserBackend string `json:"_auth_user_backend"`
		AuthUserId 		string `json:"_auth_user_id"`
	}

	SessionService struct {

	}
)

func (s *SessionService) convertAuthUser(user *django_models.AuthUser) (sessionData *SessionData) {
	sessionData.AuthUserId = string(user.Id)
	return 
}


func (s *SessionService) Get(id string) (*django_models.Session, error) {
	var djangoSession django_models.Session
	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv(DatabaseConnectionEnvironmentVariable)))), pgdialect.New())
	if err := db.NewSelect().Model(&djangoSession).Where(BunQueryString, bun.Ident("id"), id).Scan(context.Background()); err != nil {
		return &djangoSession, err
	}
	return &djangoSession, nil
}

func (s *SessionService) Create(user *django_models.AuthUser) (djangoSession *django_models.Session, err error) {
	djangoSession.SessionKey = djangoSession.CreateKey()

	sessionData := s.convertAuthUser(user)

	jsonData, err := json.Marshal(sessionData)
	if err != nil {
		return djangoSession, err
	}

	djangoSession.SessionData = djangoSession.SignObject(jsonData)

	fmt.Println(djangoSession)
	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv(DatabaseConnectionEnvironmentVariable)))), pgdialect.New())
	if _, err := db.NewInsert().Model(&djangoSession).Exec(context.Background()); err != nil {
		return djangoSession, err
	}
	
	return djangoSession, nil
}
