package comment

import (
	"context"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/router"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/entity"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
	"github.com/rs/xid"
	"net/http"
)

type CommentService interface {
	CreateComment(ctx context.Context, input CreateCommentInput) (*entity.Comment, error)
	GetCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error)
}

type Handler struct {
	logger  logger.Logger
	service CommentService
	session session.Session
}

func NewHandler(logger logger.Logger, service CommentService, session session.Session) *Handler {
	if logger == nil {
		fmt.Println("logger is nil")
		return nil
	}
	if service == nil || session == nil {
		logger.Error("invalid arguments to create comment handler")
		return nil
	}
	return &Handler{logger: logger, service: service, session: session}
}
func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.handleCreateCommentError(w, r, err)
		return
	}
	postID := r.PostFormValue("post_id")
	content := r.PostFormValue("content")
	data := h.session.Get(r, "user")
	user, err := convertionUtil.AnyToGivenType[entity.User](data)
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
	ctx := context.WithValue(r.Context(), "user", user)
	_, err = h.service.CreateComment(ctx, CreateCommentInput{
		PostID:  postID,
		Content: content,
	})
	if err != nil {
		h.handleCreateCommentError(w, r, err)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
func (h *Handler) RegisterHandlers(_router router.Router) {
	_router.Handle("/comments", router.MethodHandlers{
		http.MethodPost: h.CreateComment,
	})
}
func (h *Handler) handleCreateCommentError(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.Error("error creating comment on service")
	h.session.Put(r, "create-comment-error", err.Error())
	h.session.Put(r, "create-comment-form", r.PostForm)
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
