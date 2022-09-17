package logger_test

import (
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewZapLoggerAdapter(t *testing.T) {
	bestCases := []string{"", "qa", "local", "test", "dev", "prod"}
	worstCases := []string{",", "afasfaf", "any", "123", "afsaf"}

	for _, bestCase := range bestCases {
		zapLogger := logger.NewZapLoggerAdapter(bestCase)
		assert.NotNil(t, zapLogger)
	}
	for _, worstCase := range worstCases {
		zapLogger := logger.NewZapLoggerAdapter(worstCase)
		assert.Nil(t, zapLogger)
	}
}
