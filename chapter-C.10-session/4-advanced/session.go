package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

type SessionManager struct {
	store    sessions.Store
	valueKey string
}

func NewSessionManager(store sessions.Store) *SessionManager {
	s := new(SessionManager)
	s.valueKey = "data"
	s.store = store

	return s
}

func (s *SessionManager) Get(c echo.Context, name string) (interface{}, error) {
	session, err := s.store.Get(c.Request(), name)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, nil
	}

	if val, ok := session.Values[s.valueKey]; ok {
		return val, nil
	} else {
		return nil, nil
	}
}

func (s *SessionManager) Set(c echo.Context, name string, value interface{}) error {
	session, _ := s.store.Get(c.Request(), name)
	session.Values[s.valueKey] = value

	err := session.Save(c.Request(), c.Response())
	return err
}

func (s *SessionManager) Delete(c echo.Context, name string) error {
	session, err := s.store.Get(c.Request(), name)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	return session.Save(c.Request(), c.Response())
}
