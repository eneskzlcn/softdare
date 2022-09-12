package session

import (
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/session"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
	"net/http"
)

type GeneralSessionData struct {
	IsLoggedIn bool            `json:"is_logged_in"`
	User       UserSessionData `json:"user"`
}

/*UserSessionData keeps the values that provide user identify.
The fields can be change for usage like just with id or password.
Do not touch the name of the struct due to dependent parts. If you
change the struct name be sure to reformat related parts.

*/
type UserSessionData struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func GeneralSessionDataFromRequest(logger logger.Logger, session session.Session, r *http.Request) GeneralSessionData {
	var out GeneralSessionData
	if session.Exists(r, "user") {
		user := session.Get(r, "user")
		userData, err := convertionUtil.AnyToGivenType[UserSessionData](user)
		if err == nil {
			out.User = userData
			out.IsLoggedIn = true
		}
		logger.Debugf("Session Data Exists For User:", userData.ID)
	}
	return out
}
