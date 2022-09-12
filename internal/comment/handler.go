package comment

import (
	"context"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
	"github.com/nicolasparada/go-mux"
	"github.com/rs/xid"
	"net/http"
)

type CommentService interface {
	CreateComment(ctx context.Context, input CreateCommentInput) (*Comment, error)
	GetCommentsByPostID(ctx context.Context, postID string) ([]*Comment, error)
}
type RabbitMQClient interface {
	PushMessage(message any, queue string) error
}

type SessionProvider interface {
	Exists(r *http.Request, key string) bool
	Get(r *http.Request, key string) any
	GetString(r *http.Request, key string) string
	PopError(r *http.Request, key string) error
	Pop(r *http.Request, key string) any
	Put(r *http.Request, key string, data any)
}
type Handler struct {
	logger          logger.Logger
	service         CommentService
	sessionProvider SessionProvider
	rabbitmqClient  RabbitMQClient
}

func NewHandler(logger logger.Logger, service CommentService, sessionProvider SessionProvider, rabbitmqClient RabbitMQClient) *Handler {
	if logger == nil {
		fmt.Println("logger is nil")
		return nil
	}
	if service == nil || sessionProvider == nil || rabbitmqClient == nil {
		logger.Error("invalid arguments to create comment handler")
		return nil
	}
	return &Handler{logger: logger, service: service, sessionProvider: sessionProvider, rabbitmqClient: rabbitmqClient}
}
func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.handleCreateCommentError(w, r, err)
		return
	}
	postID := r.PostFormValue("post_id")
	content := r.PostFormValue("content")
	data := h.sessionProvider.Get(r, "user")
	user, err := convertionUtil.AnyToGivenType[User](data)
	if err != nil {
		h.logger.Error("error getting user from session")
		h.handleCreateCommentError(w, r, err)
		return
	}
	_, err = xid.FromString(user.ID)
	if err != nil {
		h.logger.Error("invalid user id for create comment")
		h.handleCreateCommentError(w, r, err)
		return
	}
	ctx := contextUtil.WithContext(r.Context(), "user", user)
	_, err = h.service.CreateComment(ctx, CreateCommentInput{
		PostID:  postID,
		Content: content,
	})
	if err != nil {
		h.handleCreateCommentError(w, r, err)
		return
	}
	err = h.rabbitmqClient.PushMessage(IncreasePostCommentCountMessage{
		PostID:         postID,
		IncreaseAmount: 1,
	}, "increase-post-comment-count")
	if err != nil {
		h.logger.Error("error publishing increase-post-comment-count message")
	}
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
func (h *Handler) RegisterHandlers(router *mux.Router) {
	router.Handle("/comments", mux.MethodHandler{
		http.MethodPost: h.CreateComment,
	})
}
func (h *Handler) handleCreateCommentError(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.Error("error creating comment on service")
	h.sessionProvider.Put(r, "create-comment-error", err.Error())
	h.sessionProvider.Put(r, "create-comment-form", r.PostForm)
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
