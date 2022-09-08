package server

import (
	"github.com/golangcollege/sessions"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type UserSessionData struct {
	Email    string
	Username string
}

type SessionData struct {
	IsLoggedIn bool
	User       UserSessionData
}

type SessionProvider struct {
	logger     *zap.SugaredLogger
	session    *sessions.Session
	sessionKey []byte
}

func (s *SessionProvider) Exists(r *http.Request, key string) bool {
	exists := s.session.Exists(r, key)
	s.logger.Debugf("SESSION EXISTS REQUEST ARRIVED FOR KEY %s AND IS EXISTS = %s", key, strconv.FormatBool(exists))
	return exists
}

func (s *SessionProvider) Get(r *http.Request, key string) any {
	s.logger.Debugf("SESSION GET REQUEST ARRIVED FOR KEY %s", key)
	return s.session.Get(r, key)
}

func (s *SessionProvider) Put(r *http.Request, key string, data interface{}) {
	s.logger.Debugf("SESSION PUT REQUEST ARRIVED FOR KEY %s and data %v", key, data)
	s.session.Put(r, key, data)
}

func (s *SessionProvider) Remove(r *http.Request, key string) {
	s.logger.Debugf("SESSION REMOVE REQUEST ARRIVED FOR KEY %s ", key)
	s.session.Remove(r, key)
}

func NewSessionProvider(logger *zap.SugaredLogger, sessionKey string) *SessionProvider {
	sessionKeyByte := []byte(sessionKey)
	session := sessions.New(sessionKeyByte)
	return &SessionProvider{session: session, logger: logger}
}
func (s *SessionProvider) Enable(handler http.Handler) http.Handler {
	s.logger.Debug("REQUEST ARRIVE FOR ENABLE THE SESSION FOR HANDLER", zap.Any("handler", handler))
	return s.session.Enable(handler)
}