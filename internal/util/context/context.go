package context

import "context"

func FromContext[T](key string, ctx context.Context) (T, bool) {
	val, exists := ctx.Value(key).(T)
	return val, exists
}

func WithContext[T](parent context.Context, key string, val T) context.Context {
	return context.WithValue(parent, key, val)
}
