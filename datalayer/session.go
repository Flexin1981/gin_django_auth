package datalayer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Flexin1981/gin_django_auth/django_models"
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

func (s *SessionService) convertAuthUser(user *django_models.AuthUser) *SessionData {
	sessionData := SessionData{}
	sessionData.AuthUserId = fmt.Sprint(user.Id)
	return &sessionData
}


func (s *SessionService) Get(sessionKey string) (*django_models.Session, error) {
	var djangoSession django_models.Session
	fmt.Println(sessionKey)
	db := GetDatabaseConnection()
	if err := db.NewSelect().Model(&djangoSession).
	Where("session_key = ?", sessionKey).
	Scan(context.Background()); err != nil {
		return &djangoSession, err
	}
	return &djangoSession, nil
}

func (s *SessionService) Create(user *django_models.AuthUser) (*django_models.Session, error) {
	djangoSession := django_models.Session{}
	djangoSession.SessionKey = djangoSession.CreateKey()

	sessionData := s.convertAuthUser(user)

	jsonData, err := json.Marshal(sessionData)
	if err != nil {
		return &djangoSession, err
	}

	djangoSession.SessionData = djangoSession.SignObject(jsonData)
	db := GetDatabaseConnection()
	if _, err := db.NewInsert().Model(&djangoSession).Exec(context.Background()); err != nil {
		return &djangoSession, err
	}
	return &djangoSession, nil
}
