package session

import (
	"errors"
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/golangcollege/sessions"
	"net/http"
	"strconv"
)

type CollegeSessionAdapter struct {
	logger     logger.Logger
	session    *sessions.Session
	sessionKey []byte
}

func NewCollegeSessionAdapter(logger logger.Logger, config config.Session) *CollegeSessionAdapter {
	if logger == nil {
		return nil
	}
	if err := validateSessionKey(config.Key); err != nil {
		logger.Errorf("Err validating session key. Error:%s ", err.Error())
		return nil
	}
	sessionKeyByte := []byte(config.Key)
	session := sessions.New(sessionKeyByte)
	return &CollegeSessionAdapter{session: session, logger: logger}
}

func (s *CollegeSessionAdapter) Exists(r *http.Request, key string) bool {
	exists := s.session.Exists(r, key)
	s.logger.Debugf("SESSION EXISTS REQUEST ARRIVED FOR KEY %s AND IS EXISTS = %t", key, strconv.FormatBool(exists))
	return exists
}

func (s *CollegeSessionAdapter) Get(r *http.Request, key string) any {
	s.logger.Debugf("SESSION GET REQUEST ARRIVED FOR KEY %s", key)
	return s.session.Get(r, key)
}

func (s *CollegeSessionAdapter) Put(r *http.Request, key string, data interface{}) {
	s.logger.Debugf("SESSION PUT REQUEST ARRIVED FOR KEY %s and data %v", key, data)
	s.session.Put(r, key, data)
}

func (s *CollegeSessionAdapter) Remove(r *http.Request, key string) {
	s.logger.Debugf("SESSION REMOVE REQUEST ARRIVED FOR KEY %s ", key)
	s.session.Remove(r, key)
}
func (s *CollegeSessionAdapter) GetString(r *http.Request, key string) string {
	s.logger.Debugf("SESSION GET STRING REQUEST ARRIVED FOR KEY %s", key)
	return s.session.GetString(r, key)
}
func (s *CollegeSessionAdapter) PopString(r *http.Request, key string) string {
	s.logger.Debugf("SESSION POP STRING REQUEST ARRIVED FOR KEY %s", key)
	return s.session.PopString(r, key)
}
func (s *CollegeSessionAdapter) Pop(r *http.Request, key string) any {
	s.logger.Debugf("SESSION POP REQUEST ARRIVED FOR KEY %s", key)
	return s.session.Pop(r, key)
}

/*PopError extracts a string that known as oops from session and converts it to oops*/
func (s *CollegeSessionAdapter) PopError(r *http.Request, key string) error {
	s.logger.Debugf("SESSION POP ERROR REQUEST ARRIVED FOR KEY %s", key)

	str := s.PopString(r, key)
	if str != "" {
		return errors.New(str)
	}
	return nil
}
func (s *CollegeSessionAdapter) Enable(handler http.Handler) http.Handler {
	s.logger.Debug("REQUEST ARRIVE FOR ENABLE THE SESSION FOR HANDLER")
	return s.session.Enable(handler)
}
