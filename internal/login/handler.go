package login

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	coreTemplate "github.com/eneskzlcn/softdare/internal/core/html/template"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/router"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/pkg"
	"html/template"
	"net/http"
	"net/url"
)

type LoginService interface {
	Login(ctx context.Context, inp LoginInput) (*User, error)
}
type Handler struct {
	logger        logger.Logger
	service       LoginService
	loginTemplate *template.Template
	session       session.Session
}

func NewHandler(logger logger.Logger, service LoginService, provider session.Session) *Handler {
	if logger == nil {
		fmt.Printf("logger can not be nil")
		return nil
	}
	if service == nil || provider == nil {
		logger.Error(ErrInvalidHandlerArgs)
		return nil
	}
	handler := Handler{logger: logger, service: service, session: provider}
	handler.init()
	return &handler
}

func (h *Handler) init() {
	gob.Register(UserSessionData{})
	h.loginTemplate = pkg.ParseTemplate("login.gohtml")
}
func (h *Handler) Show(w http.ResponseWriter, r *http.Request) {
	sessionData := sessionDataFromRequest(h.session, r, h.logger)
	h.Render(w, loginPageData{Session: sessionData}, http.StatusOK, coreTemplate.Render)
}
func (h *Handler) Render(w http.ResponseWriter, data loginPageData, statusCode int, renderFn coreTemplate.RenderFn) {
	renderFn(h.logger, w, h.loginTemplate, data, statusCode)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("LOGIN HANDLER ACCEPTED A REQUEST WITH email %s", r.PostFormValue("email"))
	if err := r.ParseForm(); err != nil {
		h.logger.Debugf("oops occurred when parsing request's form with email %s", r.PostFormValue("email"))
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	inp := LoginInput{
		Email:    r.PostFormValue("email"),
		Username: ExtractFormValue(r.Form, "username"),
	}
	user, err := h.service.Login(ctx, inp)
	if err != nil {
		h.logger.Debug("could not login")
		if errors.Is(err, ErrUserNotFound) || errors.Is(err, ErrUsernameAlreadyTaken) || errors.Is(err, ErrValidation) {
			h.Render(w, loginPageData{
				Form: r.PostForm,
				Err:  err,
			}, http.StatusBadRequest, coreTemplate.Render)
			return
		}
		http.Error(w, "could not login", http.StatusInternalServerError)
		return
	}
	h.logger.Debugf("Successfully logged in user %v", user)
	h.session.Put(r, "user", UserSessionData{ID: user.ID, Email: user.Email, Username: user.Username})
	h.logger.Debugf("Session defined for the user %v", user)

	http.Redirect(w, r, "/", http.StatusFound)
}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	h.session.Remove(r, "user")
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) RegisterHandlers(router router.Router) {
	router.Handle("/login", http.MethodGet, h.Show)
	router.Handle("/login", http.MethodPost, h.Login)
	router.Handle("/logout", http.MethodPost, h.Logout)
}

func ExtractFormValue(form url.Values, key string) *string {
	if !form.Has(key) {
		return nil
	}
	s := form.Get(key)
	return &s
}
