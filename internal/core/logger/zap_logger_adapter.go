package logger

import (
	"fmt"
	"go.uber.org/zap"
)

type ZapLoggerAdapter struct {
	logger *zap.SugaredLogger
}

func newLoggerForEnv(env string) (*zap.SugaredLogger, error) {
	if env == "" || env == "local" || env == "test" || env == "qa" || env == "dev" {
		logger, err := zap.NewDevelopment(zap.AddCaller())
		return logger.Sugar(), err
	} else {
		logger, err := zap.NewProduction(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
		return logger.Sugar(), err
	}
}

func NewZapLoggerAdapter(env string) *ZapLoggerAdapter {
	logger, err := newLoggerForEnv(env)
	if err != nil {
		fmt.Println("error logger is nil")
		return nil
	}
	zapLoggerAdapter := ZapLoggerAdapter{logger: logger}
	return &zapLoggerAdapter
}
func (z *ZapLoggerAdapter) Info(args ...interface{}) {
	z.logger.Info(args)
}
func (z *ZapLoggerAdapter) Infof(template string, args ...interface{}) {
	z.logger.Infof(template, args)
}
func (z *ZapLoggerAdapter) Error(args ...interface{}) {
	z.logger.Error(args)
}
func (z *ZapLoggerAdapter) Errorf(template string, args ...interface{}) {
	z.logger.Errorf(template, args)
}
func (z *ZapLoggerAdapter) Debug(args ...interface{}) {
	z.logger.Debug(args)
}
func (z *ZapLoggerAdapter) Debugf(template string, args ...interface{}) {
	z.logger.Debugf(template, args)
}
func (z *ZapLoggerAdapter) Fatal(args ...interface{}) {
	z.logger.Fatal(args)
}
func (z *ZapLoggerAdapter) Fatalf(template string, args ...interface{}) {
	z.logger.Fatalf(template, args)
}
func (z *ZapLoggerAdapter) Sync() error {
	return z.logger.Sync()
}
