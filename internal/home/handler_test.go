package home_test

import (
	"github.com/eneskzlcn/softdare/internal/home"
	mocks "github.com/eneskzlcn/softdare/internal/mocks/home"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewHandler(t *testing.T) {
	t.Run("given nil logger then it should return nil when NewHandler called", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockRenderer := mocks.NewMockRenderer(controller)
		mockSessionProvider := mocks.NewMockSessionProvider(controller)
		handler := home.NewHandler(nil, mockRenderer, mockSessionProvider)
		assert.Nil(t, handler)
	})
	t.Run("given nil renderer then it should return nil when NewHandler called", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockSessionProvider := mocks.NewMockSessionProvider(controller)
		handler := home.NewHandler(zap.NewExample().Sugar(), nil, mockSessionProvider)
		assert.Nil(t, handler)
	})
	t.Run("given nil session provider then it should return nil when NewHandler called", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockRenderer := mocks.NewMockRenderer(controller)
		handler := home.NewHandler(zap.NewExample().Sugar(), mockRenderer, nil)
		assert.Nil(t, handler)
	})
	t.Run("given valid args then it should return handler when newHandler called", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockRenderer := mocks.NewMockRenderer(controller)
		mockSessionProvider := mocks.NewMockSessionProvider(controller)
		handler := home.NewHandler(zap.NewExample().Sugar(), mockRenderer, mockSessionProvider)
		assert.NotNil(t, handler)
	})
}
