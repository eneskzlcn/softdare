package server_test

import (
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/internal/server"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"strings"
	"testing"
)

func TestNewSessionProvider(t *testing.T) {
	t.Run("given nil logger then it should return nil when NewSessionProvider called", func(t *testing.T) {
		session := server.NewSessionProvider(nil, config.Session{Key: strings.Repeat("a", 32)})
		assert.Nil(t, session)
	})
	t.Run("given not valid session key then it should return nil when NewSessionProvider called", func(t *testing.T) {
		session := server.NewSessionProvider(zap.NewExample().Sugar(), config.Session{Key: strings.Repeat("a", 30)})
		assert.Nil(t, session)
	})
	t.Run("given valid args  then it should return session provider when NewSessionProvider called", func(t *testing.T) {
		session := server.NewSessionProvider(zap.NewExample().Sugar(), config.Session{Key: strings.Repeat("a", 32)})
		assert.NotNil(t, session)
	})
}
