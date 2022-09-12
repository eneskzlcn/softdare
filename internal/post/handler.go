package post

import (
	"context"
	"errors"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/comment"
	coreTemplate "github.com/eneskzlcn/softdare/internal/core/html/template"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/router"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/oops"
	"github.com/eneskzlcn/softdare/internal/pkg"
	"github.com/eneskzlcn/softdare/internal/util/convertion"
	"github.com/nicolasparada/go-mux"
	"github.com/rs/xid"
	"html/template"
	"net/http"
	"time"
)

type PostService interface {
	CreatePost(ctx context.Context, in CreatePostInput) (*CreatePostResponse, error)
	GetPostByID(ctx context.Context, postID string) (*Post, error)
	IncreasePostCommentCount(ctx context.Context, postID string, increaseAmount int) (time.Time, error)
}
type CommentService interface {
	GetCommentsByPostID(ctx context.Context, postID string) ([]*comment.Comment, error)
}

type Handler struct {
	logger         logger.Logger
	service        PostService
	commentService CommentService
	template       *template.Template
	session        session.Session
}

func NewHandler(logger logger.Logger, service PostService, session session.Session, commentService CommentService) *Handler {
	if logger == nil {
		fmt.Printf("given logger is nil\n")
		return nil
	}
	if service == nil || session == nil || commentService == nil {
		logger.Error(errors.New("invalid arguments for post handler"))
		return nil
	}
	handler := Handler{logger: logger, service: service, session: session, commentService: commentService}
	handler.init()
	return &handler
}
func (h *Handler) init() {
	h.template = pkg.ParseTemplate("post.gohtml")
}
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("CREATE POST REQUEST ARRIVED")
	if err := r.ParseForm(); err != nil {
		h.logger.Error(err)
		return
	}
	data := h.session.Get(r, userContextKey)
	user, err := convertion.AnyToGivenType[User](data)
	if err != nil {
		h.logger.Errorf("can not converted session data to user struct with error %s", err.Error())
		return
	}
	ctx := context.WithValue(r.Context(), userContextKey, user)
	_, err = h.service.CreatePost(ctx, CreatePostInput{Content: r.PostFormValue("content")})
	if err != nil {
		h.logger.Error("oops creating post from server")
		h.session.Put(r, "create-post-oops", err.Error())
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func (h *Handler) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := mux.URLParam(ctx, "postID")
	_, err := xid.FromString(postID)
	if err != nil {
		h.logger.Errorf("can not parse post id: %s", postID)
		oops.RenderPage(h.logger, h.session, r, w, err, http.StatusFound, coreTemplate.Render)
		return
	}
	post, err := h.service.GetPostByID(ctx, postID)
	if err != nil {
		h.logger.Error("can not take post from service. Post ID: %s", postID)
		oops.RenderPage(h.logger, h.session, r, w, err, http.StatusFound, coreTemplate.Render)
		return
	}
	comments, err := h.commentService.GetCommentsByPostID(ctx, postID)
	if err != nil {
		h.logger.Errorf("can not take comments from service. Post ID: %s", postID)
		oops.RenderPage(h.logger, h.session, r, w, err, http.StatusFound, coreTemplate.Render)
		return
	}
	formattedPost := FormatPost(post)
	h.logger.Debugf("FORMATTED POST: %v", formattedPost)
	formattedComments := FormatComments(comments)
	sessionData := sessionDataFromRequest(h.session, r, h.logger)
	h.Render(w, postData{Post: formattedPost, Session: sessionData, Comments: formattedComments}, http.StatusFound, coreTemplate.Render)
}
func (h *Handler) Render(w http.ResponseWriter, data postData, statusCode int, renderFn coreTemplate.RenderFn) {
	h.logger.Debug("RENDERING THE POST TEMPLATE")
	renderFn(h.logger, w, h.template, data, statusCode)
}
func (h *Handler) RegisterHandlers(router router.Router) {
	router.Handle("/posts", http.MethodPost, h.CreatePost)
	router.Handle("/posts/{postID}", http.MethodGet, h.Show)
}
