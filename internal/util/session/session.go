package session

import (
	"fmt"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
	"net/http"
)

type GeneralSessionData struct {
	IsLoggedIn bool            `json:"is_logged_in"`
	User       UserSessionData `json:"user"`
}
type UserSessionData struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
type SessionProvider interface {
	Exists(r *http.Request, key string) bool
	Get(r *http.Request, key string) any
}

func GeneralSessionDataFromRequest(session SessionProvider, r *http.Request) GeneralSessionData {
	var out GeneralSessionData
	if session.Exists(r, "user") {
		user := session.Get(r, "user")
		userData, err := convertionUtil.AnyToGivenType[UserSessionData](user)
		if err == nil {
			out.User = userData
			out.IsLoggedIn = true
		}
		fmt.Printf("Session data exist for the user. Session data:%v\n", out)
	}
	return out
}
