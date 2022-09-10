package server_test

import (
	mocks "github.com/eneskzlcn/softdare/internal/mocks/server"
	server "github.com/eneskzlcn/softdare/internal/server"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewHandler(t *testing.T) {
	type testCase struct {
		logger          *zap.SugaredLogger
		routeHandlers   []server.RouteHandler
		sessionProvider server.Session
		scenario        string
		expectedObject  *server.Handler
		expectedError   error
	}
	exLogger := zap.NewExample().Sugar()
	controller := gomock.NewController(t)
	exSessionProvider := mocks.NewMockSession(controller)
	exSessionProvider.EXPECT().Enable(gomock.Any()).Return(nil)
	handlerCreateScenarios := []testCase{
		{
			logger:          exLogger,
			routeHandlers:   nil,
			sessionProvider: nil,
			scenario:        "test given empty session provider then it should return nil with oops when new handler called",
			expectedError:   server.ErrSessionProviderNil,
		},
		{
			logger:          nil,
			routeHandlers:   nil,
			sessionProvider: exSessionProvider,
			scenario:        "test given empty logger then it should return nil with oops when new handler called",
			expectedError:   server.ErrLoggerNil,
		},
		{
			logger:          exLogger,
			routeHandlers:   []server.RouteHandler{nil, nil, nil},
			sessionProvider: exSessionProvider,
			scenario:        "test given a nil route handler then it should return nil with oops when new handler called",
			expectedError:   server.ErrGivenRouteHandlerNil,
		},
		{
			logger:          exLogger,
			routeHandlers:   nil,
			sessionProvider: exSessionProvider,
			scenario:        "test given valid arguments then it should return handler without oops when new handler called",
			expectedError:   nil,
		},
	}
	for _, testCase := range handlerCreateScenarios {
		_, err := server.NewHandler(testCase.logger, testCase.routeHandlers, testCase.sessionProvider)
		assert.Equal(t, err, testCase.expectedError)
	}
}
