package logger

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Fatal(args ...interface{})
	Errorf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Debugf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Sync() error
}
