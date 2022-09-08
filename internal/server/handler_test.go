package server_test

import (
	"github.com/eneskzlcn/softdare/internal/config"
	server2 "github.com/eneskzlcn/softdare/internal/server"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewHandler(t *testing.T) {
	type testCase struct {
		logger          *zap.SugaredLogger
		routeHandlers   []server2.RouteHandler
		sessionProvider server2.Session
		scenario        string
		expectedObject  *server2.Handler
		expectedError   error
	}
	exLogger := zap.NewExample().Sugar()
	exSessionProvider := server2.NewSessionProvider(zap.NewExample().Sugar(), config.Session{Key: "test"})
	handlerCreateScenarios := []testCase{
		{
			logger:          exLogger,
			routeHandlers:   nil,
			sessionProvider: nil,
			scenario:        "test given empty session provider then it should return nil with error when new handler called",
			expectedError:   server2.ErrSessionProviderNil,
		},
		{
			logger:          nil,
			routeHandlers:   nil,
			sessionProvider: exSessionProvider,
			scenario:        "test given empty logger then it should return nil with error when new handler called",
			expectedError:   server2.ErrLoggerNil,
		},
		{
			logger:          exLogger,
			routeHandlers:   []server2.RouteHandler{nil, nil, nil},
			sessionProvider: exSessionProvider,
			scenario:        "test given a nil route handler then it should return nil with error when new handler called",
			expectedError:   server2.ErrGivenRouteHandlerNil,
		},
		{
			logger:          exLogger,
			routeHandlers:   nil,
			sessionProvider: exSessionProvider,
			scenario:        "test given valid arguments then it should return handler without error when new handler called",
			expectedError:   nil,
		},
	}
	for _, testCase := range handlerCreateScenarios {
		_, err := server2.NewHandler(testCase.logger, testCase.routeHandlers, testCase.sessionProvider)
		assert.Equal(t, err, testCase.expectedError)
	}
}
