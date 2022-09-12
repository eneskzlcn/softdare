package oops

import (
	coreTemplate "github.com/eneskzlcn/softdare/internal/core/html/template"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/pkg"
	sessionUtil "github.com/eneskzlcn/softdare/internal/util/session"
	"net/http"
)

const DomainName = "oops"

type ErrData struct {
	Err     error
	Session SessionData
}
type SessionData struct {
	IsLoggedIn bool `json:"is_logged_in"`
	User       sessionUtil.UserSessionData
}

func RenderPage(logger logger.Logger, session session.Session, r *http.Request, w http.ResponseWriter, data error, statusCode int, renderFn coreTemplate.RenderFn) {
	logger.Debugf("Rendering the ooops page.")
	tmpl := pkg.ParseTemplate("oops.gohtml")
	generalSessionData := sessionUtil.GeneralSessionDataFromRequest(logger, session, r)
	sessionData := SessionData{
		IsLoggedIn: generalSessionData.IsLoggedIn,
		User:       generalSessionData.User,
	}
	renderFn(logger, w, tmpl, ErrData{Err: data, Session: sessionData}, statusCode)
}
