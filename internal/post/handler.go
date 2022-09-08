package post

import (
	"context"
	"fmt"
	mux_router "github.com/eneskzlcn/mux-router"
	contextUtil "github.com/eneskzlcn/softdare/internal/context"
	"go.uber.org/zap"
	"net/http"
)

type PostService interface {
	CreatePost(ctx context.Context, in CreatePostInput) (*CreatePostResponse, error)
}
type Renderer interface {
}
type SessionProvider interface {
	Get(r *http.Request, key string) any
}
type Handler struct {
	logger          *zap.SugaredLogger
	service         PostService
	renderer        Renderer
	sessionProvider SessionProvider
}

func NewHandler(logger *zap.SugaredLogger, service PostService, renderer Renderer, sessionProvider SessionProvider) *Handler {
	if logger == nil {
		fmt.Printf("given logger is nil\n")
		return nil
	}
	if service == nil {
		logger.Error(ErrPostServiceNil)
		return nil
	}
	if renderer == nil {
		logger.Error(ErrRendererNil)
		return nil
	}
	if sessionProvider == nil {
		logger.Error(ErrSessionProviderNil)
		return nil
	}
	handler := Handler{logger: logger, service: service, renderer: renderer}
	if err := handler.init(); err != nil {
		logger.Error(err)
		return nil
	}
	return &handler
}
func (h *Handler) init() error {
	return nil
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("CREATE POST REQUEST ARRIVED")
	if err := r.ParseForm(); err != nil {
		h.logger.Error(err)
		return
	}
	user, exists := h.sessionProvider.Get(r, userContextKey).(User)
	if !exists {
		h.logger.Error("can not get user from session provider", zap.Any("user", user))
		return
	}
	ctx := contextUtil.ContextWith[User](r.Context(), userContextKey, user)
	_, err := h.service.CreatePost(ctx, CreatePostInput{Content: r.PostFormValue("content")})
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	http.Redirect(w, r, "/home", http.StatusFound)
}
func (h *Handler) RegisterHandlers(router *mux_router.Router) {
	router.HandleFunc(http.MethodPost, "/posts", h.CreatePost)
}
