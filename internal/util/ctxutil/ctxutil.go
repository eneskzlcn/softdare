package ctxutil

import "context"

/*FromContext generic function takes a key and context
and tries to cast the value of the key to given generic type T.
if value not exists or can not get converted it will return nil with false.
Otherwise, it returns the value in type T with true.
*/
func FromContext[T any](key any, ctx context.Context) (T, bool) {
	val, exists := ctx.Value(key).(T)
	return val, exists
}
