package post_test

//
//import (
//	mux_router "github.com/eneskzlcn/mux-router"
//	mocks "github.com/eneskzlcn/softdare/internal/mocks/post"
//	"github.com/eneskzlcn/softdare/internal/post"
//	"github.com/golang/mock/gomock"
//	"github.com/rs/xid"
//	"github.com/stretchr/testify/assert"
//	"go.uber.org/zap"
//	"net/http"
//	"net/url"
//	"testing"
//)
//
//func TestNewHandler(t *testing.T) {
//
//	controller := gomock.NewController(t)
//	mockPostService := mocks.NewMockPostService(controller)
//	mockSessionProvider := mocks.NewMockSessionProvider(controller)
//	mockRenderer := mocks.NewMockRenderer(controller)
//	commentService := mocks.NewMockCommentService(controller)
//	logger := zap.NewExample().Sugar()
//
//	t.Run("given nil logger then it should return nil when NewHandler called", func(t *testing.T) {
//		handler := post.NewHandler(nil, mockPostService, mockSessionProvider, mockRenderer, commentService)
//		assert.Nil(t, handler)
//	})
//	t.Run("given nil post service then it should return nil when NewHandler called", func(t *testing.T) {
//		handler := post.NewHandler(logger, nil, mockSessionProvider, mockRenderer, commentService)
//		assert.Nil(t, handler)
//	})
//	t.Run("given nil session provider then it should return nil when NewHandler called", func(t *testing.T) {
//		handler := post.NewHandler(logger, mockPostService, nil, mockRenderer, commentService)
//		assert.Nil(t, handler)
//	})
//	t.Run("given nil renderer then it should return nil when NewHandler called", func(t *testing.T) {
//		handler := post.NewHandler(logger, mockPostService, mockSessionProvider, nil, commentService)
//		assert.Nil(t, handler)
//	})
//	t.Run("given nil comment service then it should return nil when NewHandler called", func(t *testing.T) {
//		handler := post.NewHandler(logger, mockPostService, mockSessionProvider, mockRenderer, nil)
//		assert.Nil(t, handler)
//	})
//	t.Run("given valid args then it should return handler when NewHandler called", func(t *testing.T) {
//		handler := post.NewHandler(logger, mockPostService, mockSessionProvider, mockRenderer, commentService)
//		assert.NotNil(t, handler)
//	})
//}
//func TestHandler_CreatePost(t *testing.T) {
//	controller := gomock.NewController(t)
//	mockPostService := mocks.NewMockPostService(controller)
//	mockSessionProvider := mocks.NewMockSessionProvider(controller)
//	logger := zap.NewExample().Sugar()
//	handler := post.NewHandler(logger, mockPostService, mockSessionProvider, nil, nil)
//
//	router := mux_router.New()
//	router.HandleFunc(http.MethodPost, "/posts", handler.CreatePost)
//	router.HandleFunc(http.MethodGet, "/", func(writer http.ResponseWriter, request *http.Request) {
//
//	})
//	server := http.Server{Addr: ":4300", Handler: router}
//
//	go server.ListenAndServe()
//
//	t.Run("given valid content", func(t *testing.T) {
//		content := "valid content"
//		req, err := http.NewRequest(http.MethodPost, "http://localhost:4300/posts", nil)
//		assert.Nil(t, err)
//		req.PostForm = make(url.Values)
//		req.PostForm.Set("content", content)
//		user := post.User{
//			ID:       xid.New().String(),
//			Email:    "valid@sm.com",
//			Username: "valids",
//		}
//		mockSessionProvider.EXPECT().Get(gomock.Any(), "user").Return(user)
//		mockPostService.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Return(nil, nil)
//		resp, err := http.DefaultClient.Do(req)
//		assert.Nil(t, err)
//		assert.Equal(t, http.StatusOK, resp.StatusCode)
//	})
//	t.Run("given not valid content then there should be oops", func(t *testing.T) {
//		content := ""
//		req, err := http.NewRequest(http.MethodPost, "http://localhost:4300/posts", nil)
//		assert.Nil(t, err)
//		req.PostForm = make(url.Values)
//		req.PostForm.Set("content", content)
//		user := post.User{
//			ID:       xid.New().String(),
//			Email:    "valid@sm.com",
//			Username: "valids",
//		}
//		mockSessionProvider.EXPECT().Get(gomock.Any(), "user").Return(user)
//		mockPostService.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Return(nil, err)
//		resp, err := http.DefaultClient.Do(req)
//		assert.Nil(t, err)
//		assert.Equal(t, http.StatusOK, resp.StatusCode)
//	})
//}
