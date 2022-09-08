package session

import "encoding/json"

func SessionDataFromAny[T](data any) (T, error) {
	var sessionData T
	bytes, err := json.Marshal(data)
	if err != nil {
		return sessionData, err
	}
	err = json.Unmarshal(bytes, &sessionData)
	return sessionData, err
}
