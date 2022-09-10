package home_test

import (
	mux_router "github.com/eneskzlcn/mux-router"
	"github.com/eneskzlcn/softdare/internal/home"
	mocks "github.com/eneskzlcn/softdare/internal/mocks/home"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"testing"
)

func TestNewHandler(t *testing.T) {
	controller := gomock.NewController(t)
	mockRenderer := mocks.NewMockRenderer(controller)
	mockSessionProvider := mocks.NewMockSessionProvider(controller)
	mockHomeService := mocks.NewMockHomeService(controller)
	t.Run("given nil logger then it should return nil when NewHandler called", func(t *testing.T) {
		handler := home.NewHandler(nil, mockRenderer, mockSessionProvider, mockHomeService)
		assert.Nil(t, handler)
	})
	t.Run("given nil renderer then it should return nil when NewHandler called", func(t *testing.T) {
		handler := home.NewHandler(zap.NewExample().Sugar(), nil, mockSessionProvider, mockHomeService)
		assert.Nil(t, handler)
	})
	t.Run("given nil session provider then it should return nil when NewHandler called", func(t *testing.T) {
		handler := home.NewHandler(zap.NewExample().Sugar(), mockRenderer, nil, mockHomeService)
		assert.Nil(t, handler)
	})
	t.Run("given nil home service then it should return nil when newHandler called", func(t *testing.T) {
		handler := home.NewHandler(zap.NewExample().Sugar(), mockRenderer, mockSessionProvider, nil)
		assert.Nil(t, handler)
	})
	t.Run("given valid args then it should return handler when newHandler called", func(t *testing.T) {
		handler := home.NewHandler(zap.NewExample().Sugar(), mockRenderer, mockSessionProvider, mockHomeService)
		assert.NotNil(t, handler)
	})
}

func TestHandler_Show(t *testing.T) {
	controller := gomock.NewController(t)
	mockSessionProvider := mocks.NewMockSessionProvider(controller)
	mockRenderer := mocks.NewMockRenderer(controller)
	logger := zap.NewExample().Sugar()
	handler := home.NewHandler(logger, mockRenderer, mockSessionProvider, nil)
	router := mux_router.New()
	router.HandleFunc(http.MethodGet, "/", handler.Show)
	mockRenderer.EXPECT().RenderTemplate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
	server := http.Server{
		Addr:    ":4300",
		Handler: router,
	}
	go server.ListenAndServe()
	type UserSessionData struct {
		Email    string
		Username string
	}
	mockSessionProvider.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(true)
	mockSessionProvider.EXPECT().Get(gomock.Any(), "user").Return(UserSessionData{Email: "a@s.com", Username: "asdad"})

	req, err := http.NewRequest(http.MethodGet, "http://localhost:4300", nil)
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
