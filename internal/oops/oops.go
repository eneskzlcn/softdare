package oops

import (
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/pkg"
	sessionUtil "github.com/eneskzlcn/softdare/internal/util/session"
	"html/template"
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
type Renderer interface {
	RenderTemplate(w http.ResponseWriter, template *template.Template, data any, statusCode int)
}
type SessionProvider interface {
	Exists(r *http.Request, key string) bool
	Get(r *http.Request, key string) any
}

func RenderPage(renderer Renderer, logger logger.Logger, sessionProvider SessionProvider, r *http.Request, w http.ResponseWriter, data error, statusCode int) {
	logger.Debugf("Rendering the ooops page.")
	tmpl := pkg.ParseTemplate("oops.gohtml")
	generalSessionData := sessionUtil.GeneralSessionDataFromRequest(sessionProvider, r)
	session := SessionData{
		IsLoggedIn: generalSessionData.IsLoggedIn,
		User:       generalSessionData.User,
	}
	renderer.RenderTemplate(w, tmpl, ErrData{Err: data, Session: session}, statusCode)
}
