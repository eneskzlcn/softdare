package home

import "encoding/json"

const DomainName = "home"

type UserSessionData struct {
	Email    string
	Username string
}

type SessionData struct {
	IsLoggedIn bool
	User       UserSessionData
}

type homeData struct {
	Session SessionData
}

func SessionDataFromAny(data any) (UserSessionData, error) {
	var sessionData UserSessionData
	bytes, err := json.Marshal(data)
	if err != nil {
		return sessionData, err
	}
	err = json.Unmarshal(bytes, &sessionData)
	return sessionData, err
}
