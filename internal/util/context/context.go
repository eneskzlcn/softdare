package context

import "context"

func FromContext[T any](key any, ctx context.Context) (T, bool) {
	val, exists := ctx.Value(key).(T)
	return val, exists
}
