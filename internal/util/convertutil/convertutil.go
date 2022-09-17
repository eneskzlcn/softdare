package convertutil

import "encoding/json"

/*AnyTo generic function takes a
data and a generic type T. Then tries to convert the given data
to type of T. It returns nil with error if it cannot convert. Otherwise,
it returns the data in converted format.*/
func AnyTo[T any](data any) (T, error) {
	var to T
	bytes, err := json.Marshal(data)
	if err != nil {
		return to, err
	}
	err = json.Unmarshal(bytes, &to)
	return to, err
}
