package logger

import "go.uber.org/zap"

func NewLoggerForEnv(env string) (*zap.SugaredLogger, error) {
	if env == "" || env == "local" || env == "test" || env == "qa" || env == "dev" {
		logger, err := zap.NewDevelopment(zap.AddCaller())
		return logger.Sugar(), err
	} else {
		logger, err := zap.NewProduction(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
		return logger.Sugar(), err
	}
}
