package session

import (
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/entity"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
	"net/http"
)

func GeneralSessionDataFromRequest(logger logger.Logger, session session.Session, r *http.Request) (isLoggedIn bool, user entity.UserSessionData) {
	if session.Exists(r, "user") {
		data := session.Get(r, "user")
		userData, err := convertionUtil.AnyToGivenType[entity.UserSessionData](data)
		if err == nil {
			user = userData
			isLoggedIn = true
		}
		logger.Debugf("Session Data Exists For User:", userData.ID)
	}
	return
}
