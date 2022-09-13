package login_test

//
//import (
//	mux_router "github.com/eneskzlcn/mux-router"
//	"github.com/eneskzlcn/softdare/internal/login"
//	mocks "github.com/eneskzlcn/softdare/internal/mocks/login"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	"go.uber.org/zap"
//	"net/http"
//	"net/url"
//	"strings"
//	"testing"
//	"time"
//)
//
//func TestNewHandler(t *testing.T) {
//	//t.Run("given nil logger then it should return nil when NewHandler called", func(t *testing.T) {
//	//	controller := gomock.NewController(t)
//	//	mockSessionProvider := mocks.NewMockSessionProvider(controller)
//	//	mockRenderer := mocks.NewMockRenderer(controller)
//	//	mockLoginService := mocks.NewMockLoginService(controller)
//	//	handler := login.NewHandler(nil, mockLoginService, mockRenderer, mockSessionProvider)
//	//	assert.Nil(t, handler)
//	//})
//	//t.Run("given nil renderer then it should return nil when NewHandler called", func(t *testing.T) {
//	//	controller := gomock.NewController(t)
//	//	mockSessionProvider := mocks.NewMockSessionProvider(controller)
//	//	mockLoginService := mocks.NewMockLoginService(controller)
//	//	handler := login.NewHandler(zap.NewExample().Sugar(), mockLoginService, nil, mockSessionProvider)
//	//	assert.Nil(t, handler)
//	//})
//	//t.Run("given nil login service then it should return nil when NewHandler called", func(t *testing.T) {
//	//	controller := gomock.NewController(t)
//	//	mockSessionProvider := mocks.NewMockSessionProvider(controller)
//	//	mockRenderer := mocks.NewMockRenderer(controller)
//	//	handler := login.NewHandler(zap.NewExample().Sugar(), nil, mockRenderer, mockSessionProvider)
//	//	assert.Nil(t, handler)
//	//})
//	//t.Run("given nil session provider then it should return nil when NewHandler called", func(t *testing.T) {
//	//	controller := gomock.NewController(t)
//	//	mockRenderer := mocks.NewMockRenderer(controller)
//	//	mockLoginService := mocks.NewMockLoginService(controller)
//	//	handler := login.NewHandler(zap.NewExample().Sugar(), mockLoginService, mockRenderer, nil)
//	//	assert.Nil(t, handler)
//	//})
//	//t.Run("given valid args then it should return handler when NewHandler called", func(t *testing.T) {
//	//	controller := gomock.NewController(t)
//	//	mockSessionProvider := mocks.NewMockSessionProvider(controller)
//	//	mockRenderer := mocks.NewMockRenderer(controller)
//	//	mockLoginService := mocks.NewMockLoginService(controller)
//	//	handler := login.NewHandler(zap.NewExample().Sugar(), mockLoginService, mockRenderer, mockSessionProvider)
//	//	assert.NotNil(t, handler)
//	//})
//}
//func TestHandler_Show(t *testing.T) {
//	controller := gomock.NewController(t)
//	mockSessionProvider := mocks.NewMockSessionProvider(controller)
//	mockRenderer := mocks.NewMockRenderer(controller)
//	mockLoginService := mocks.NewMockLoginService(controller)
//	handler := login.NewHandler(zap.NewExample().Sugar(), mockLoginService, mockRenderer, mockSessionProvider)
//
//	mockRenderer.EXPECT().RenderTemplate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1)
//	router := mux_router.New()
//	router.HandleFunc(http.MethodGet, "/login", handler.Show)
//	server := http.Server{Handler: router, Addr: ":4200"}
//
//	go server.ListenAndServe()
//
//	req, err := http.NewRequest(http.MethodGet, "http://localhost:4200/login", nil)
//	assert.Nil(t, err)
//	resp, err := http.DefaultClient.Do(req)
//	assert.Nil(t, err)
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//}
//
//func TestHandler_Login(t *testing.T) {
//	controller := gomock.NewController(t)
//	mockSessionProvider := mocks.NewMockSessionProvider(controller)
//	mockRenderer := mocks.NewMockRenderer(controller)
//	mockLoginService := mocks.NewMockLoginService(controller)
//	handler := login.NewHandler(zap.NewExample().Sugar(), mockLoginService, mockRenderer, mockSessionProvider)
//
//	testLoginEmail := "exampleEmail@example.com"
//	testLoginUsername := "example"
//
//	mockRenderer.EXPECT().RenderTemplate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1)
//	mockSessionProvider.EXPECT().Put(gomock.Any(), "user", login.UserSessionData{
//		Email:    testLoginEmail,
//		Username: testLoginUsername,
//	})
//	router := mux_router.New()
//	router.HandleFunc(http.MethodPost, "/login", handler.Login)
//	//it redirects to home after logic ends, so I add a mock render page below to avoid oops.
//	router.HandleFunc(http.MethodGet, "/", handler.Show)
//	server := http.Server{Handler: router, Addr: ":4200"}
//	go server.ListenAndServe()
//
//	mockLoginService.EXPECT().Login(gomock.Any(), gomock.Any()).Return(&login.User{
//		ID:        strings.Repeat("12ase", 4),
//		Email:     testLoginEmail,
//		Username:  testLoginUsername,
//		CreatedAt: time.Now(),
//		UpdatedAt: time.Now(),
//	}, nil)
//
//	req, err := http.NewRequest(http.MethodPost, "http://localhost:4200/login", nil)
//	req.PostForm = make(url.Values)
//	req.Form = make(url.Values)
//	req.PostForm.Set("email", testLoginEmail)
//	req.Form.Set("username", testLoginUsername)
//	assert.Nil(t, err)
//	resp, err := http.DefaultClient.Do(req)
//	assert.Nil(t, err)
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//}
//func TestHandler_Logout(t *testing.T) {
//	controller := gomock.NewController(t)
//	mockSessionProvider := mocks.NewMockSessionProvider(controller)
//	mockRenderer := mocks.NewMockRenderer(controller)
//	mockLoginService := mocks.NewMockLoginService(controller)
//	handler := login.NewHandler(zap.NewExample().Sugar(), mockLoginService, mockRenderer, mockSessionProvider)
//
//	mockRenderer.EXPECT().RenderTemplate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1)
//	router := mux_router.New()
//
//	mockSessionProvider.EXPECT().Remove(gomock.Any(), "user").MinTimes(1)
//
//	router.HandleFunc(http.MethodPost, "/logout", handler.Logout)
//	router.HandleFunc(http.MethodGet, "/", handler.Show)
//	server := http.Server{Handler: router, Addr: ":4200"}
//
//	go server.ListenAndServe()
//
//	req, err := http.NewRequest(http.MethodPost, "http://localhost:4200/logout", nil)
//	assert.Nil(t, err)
//	resp, err := http.DefaultClient.Do(req)
//	assert.Nil(t, err)
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//}
