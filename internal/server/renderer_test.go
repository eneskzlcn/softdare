package server_test

import (
	"github.com/eneskzlcn/softdare/internal/server"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewRenderer(t *testing.T) {
	t.Run("given nil logger then it should return nil when NewRenderer called", func(t *testing.T) {
		renderer := server.NewRenderer(nil)
		assert.Nil(t, renderer)
	})
	t.Run("given valid args then it should return renderer when NewRenderer called", func(t *testing.T) {
		renderer := server.NewRenderer(zap.NewExample().Sugar())
		assert.NotNil(t, renderer)
	})
}
