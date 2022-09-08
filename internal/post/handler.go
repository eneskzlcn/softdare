package post

import (
	"context"
	"fmt"
	mux_router "github.com/eneskzlcn/mux-router"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
	"go.uber.org/zap"
	"net/http"
)

type PostService interface {
	CreatePost(ctx context.Context, in CreatePostInput) (*CreatePostResponse, error)
}
type SessionProvider interface {
	Get(r *http.Request, key string) any
}
type Handler struct {
	logger          *zap.SugaredLogger
	service         PostService
	sessionProvider SessionProvider
}

func NewHandler(logger *zap.SugaredLogger, service PostService, sessionProvider SessionProvider) *Handler {
	if logger == nil {
		fmt.Printf("given logger is nil\n")
		return nil
	}
	if service == nil {
		logger.Error(ErrPostServiceNil)
		return nil
	}

	if sessionProvider == nil {
		logger.Error(ErrSessionProviderNil)
		return nil
	}
	handler := Handler{logger: logger, service: service, sessionProvider: sessionProvider}
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
	data := h.sessionProvider.Get(r, userContextKey)
	user, err := convertionUtil.AnyToGivenType[User](data)
	if err != nil {
		h.logger.Error("can not converted session data to user struct", zap.Error(err))
		return
	}
	ctx := contextUtil.WithContext[User](r.Context(), userContextKey, user)
	_, err = h.service.CreatePost(ctx, CreatePostInput{Content: r.PostFormValue("content")})
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func (h *Handler) RegisterHandlers(router *mux_router.Router) {
	router.HandleFunc(http.MethodPost, "/posts", h.CreatePost)
}
