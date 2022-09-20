package web_test

//
//import (
//	"github.com/eneskzlcn/softdare/internal/entity"
//	mocks "github.com/eneskzlcn/softdare/internal/mocks/web"
//	"github.com/eneskzlcn/softdare/internal/web"
//	"github.com/stretchr/testify/assert"
//
//	"github.com/golang/mock/gomock"
//	"testing"
//)
//
//func TestNewHandler(t *testing.T) {
//	controller := gomock.NewController(t)
//	mockService := mocks.NewMockService(controller)
//	mockRenderer := mocks.NewMockRenderer(controller)
//	mockLogger := mocks.NewMockLogger(controller)
//	mockSession := mocks.NewMockSession(controller)
//	t.Run("given nil logger then it should return nil", func(t *testing.T) {
//		handler := web.NewHandler(nil, mockSession, mockService, mockRenderer)
//		assert.Nil(t, handler)
//	})
//	t.Run("given nil session then it should return nil", func(t *testing.T) {
//		mockLogger.EXPECT().Error(entity.InvalidConstructorArguments).Times(1)
//		handler := web.NewHandler(mockLogger, nil, mockService, mockRenderer)
//		assert.Nil(t, handler)
//	})
//	t.Run("given nil service then it should return nil", func(t *testing.T) {
//		mockLogger.EXPECT().Error(entity.InvalidConstructorArguments).Times(1)
//		handler := web.NewHandler(mockLogger, mockSession, nil, mockRenderer)
//		assert.Nil(t, handler)
//	})
//	t.Run("given nil renderer then it should return nil", func(t *testing.T) {
//		mockLogger.EXPECT().Error(entity.InvalidConstructorArguments).Times(1)
//		handler := web.NewHandler(mockLogger, mockSession, mockService, nil)
//		assert.Nil(t, handler)
//	})
//	t.Run("given valid arguments then it should return handler", func(t *testing.T) {
//		mockSession.EXPECT().Enable(gomock.Any()).Times(1)
//		handler := web.NewHandler(mockLogger, mockSession, mockService, mockRenderer)
//		assert.NotNil(t, handler)
//	})
//}
