package server_test

import (
	config "github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/internal/server"
	"github.com/eneskzlcn/softdare/internal/server/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewServer(t *testing.T) {
	t.Run("given nil logger then it should return nil when new called", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockRootHandler := mocks.NewMockRootHandler(controller)
		conf := config.Server{Address: ":4000"}
		serv := server.New(conf, mockRootHandler, nil)
		assert.Nil(t, serv)
	})
	t.Run("given nil rootHandler then it should return nil when new called", func(t *testing.T) {
		logger := zap.NewExample().Sugar()
		conf := config.Server{Address: ":4000"}
		serv := server.New(conf, nil, logger)
		assert.Nil(t, serv)
	})
	t.Run("given not valid config then it should return nil when new called", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockRootHandler := mocks.NewMockRootHandler(controller)
		wrongConfigs := []config.Server{
			{Address: ""},
			{Address: ":1asssf"},
			{Address: ":asdad"},
			{Address: "4000:"},
		}
		for _, conf := range wrongConfigs {
			serv := server.New(conf, mockRootHandler, zap.NewExample().Sugar())
			assert.Nil(t, serv)
		}
	})
	t.Run("given valid arguments then it should return not nil when new called", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockRootHandler := mocks.NewMockRootHandler(controller)
		conf := config.Server{Address: ":4000"}
		serv := server.New(conf, mockRootHandler, zap.NewExample().Sugar())
		assert.NotNil(t, serv)

	})
}
