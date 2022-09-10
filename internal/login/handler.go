package login

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	muxRouter "github.com/eneskzlcn/mux-router"
	"github.com/eneskzlcn/softdare/internal/pkg"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"net/url"
)

type Renderer interface {
	RenderTemplate(w http.ResponseWriter, template *template.Template, data any, statusCode int)
}
type SessionProvider interface {
	Put(r *http.Request, key string, data any)
	Remove(r *http.Request, key string)
}
type LoginService interface {
	Login(ctx context.Context, inp LoginInput) (*User, error)
}
type Handler struct {
	logger          *zap.SugaredLogger
	service         LoginService
	loginTemplate   *template.Template
	renderer        Renderer
	sessionProvider SessionProvider
}

func NewHandler(logger *zap.SugaredLogger, service LoginService, renderer Renderer, provider SessionProvider) *Handler {
	if logger == nil {
		fmt.Printf("logger can not be nil")
		return nil
	}
	if service == nil || renderer == nil || provider == nil {
		logger.Error(ErrInvalidHandlerArgs)
		return nil
	}
	handler := Handler{logger: logger, service: service, renderer: renderer, sessionProvider: provider}
	if err := handler.init(); err != nil {
		logger.Error(err)
		return nil
	}
	return &handler
}

func (h *Handler) init() error {
	template, err := pkg.ParseTemplate(DomainName)
	if err != nil {
		return err
	}
	h.loginTemplate = template
	gob.Register(UserSessionData{})
	return nil
}
func (h *Handler) Show(w http.ResponseWriter, r *http.Request) {
	h.Render(w, loginPageData{}, http.StatusOK)
}
func (h *Handler) Render(w http.ResponseWriter, data loginPageData, statusCode int) {
	h.renderer.RenderTemplate(w, h.loginTemplate, data, statusCode)
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
		Email: r.PostFormValue("email"),

		Username: ExtractFormValue(r.Form, "username"),
	}
	user, err := h.service.Login(ctx, inp)
	if err != nil {
		h.logger.Debug("could not login", zap.Error(err))
		if errors.Is(err, ErrUserNotFound) || errors.Is(err, ErrUsernameAlreadyTaken) || errors.Is(err, ErrValidation) {
			h.Render(w, loginPageData{
				Form: r.PostForm,
				Err:  err,
			}, http.StatusBadRequest)
			return
		}
		http.Error(w, "could not login", http.StatusInternalServerError)
		return
	}
	h.logger.Debugf("Successfully logged in user %v", user)
	h.sessionProvider.Put(r, "user", UserSessionData{ID: user.ID, Email: user.Email, Username: user.Username})
	h.logger.Debugf("Session defined for the user %v", user)

	http.Redirect(w, r, "/", http.StatusFound)
}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	h.sessionProvider.Remove(r, "user")
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) RegisterHandlers(router *muxRouter.Router) {
	router.HandleFunc(http.MethodGet, "/login", h.Show)
	router.HandleFunc(http.MethodPost, "/login", h.Login)
	router.HandleFunc(http.MethodPost, "/logout", h.Logout)
}

func ExtractFormValue(form url.Values, key string) *string {
	if !form.Has(key) {
		return nil
	}
	s := form.Get(key)
	return &s
}
