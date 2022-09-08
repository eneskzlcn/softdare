package convertion

import "encoding/json"

func AnyToGivenType[T any](data any) (T, error) {
	var sessionData T
	bytes, err := json.Marshal(data)
	if err != nil {
		return sessionData, err
	}
	err = json.Unmarshal(bytes, &sessionData)
	return sessionData, err
}
