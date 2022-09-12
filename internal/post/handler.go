package post

import (
	"context"
	"errors"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/comment"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/oops"
	"github.com/eneskzlcn/softdare/internal/pkg"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
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
type Renderer interface {
	RenderTemplate(w http.ResponseWriter, template *template.Template, data any, statusCode int)
}
type SessionProvider interface {
	Get(r *http.Request, key string) any
	Put(r *http.Request, key string, data interface{})
	Exists(r *http.Request, key string) bool
	PopError(r *http.Request, key string) error
}
type Handler struct {
	logger          logger.Logger
	service         PostService
	commentService  CommentService
	template        *template.Template
	sessionProvider SessionProvider
	renderer        Renderer
}

func NewHandler(logger logger.Logger, service PostService, sessionProvider SessionProvider, renderer Renderer, commentService CommentService) *Handler {
	if logger == nil {
		fmt.Printf("given logger is nil\n")
		return nil
	}
	if service == nil || sessionProvider == nil || renderer == nil || commentService == nil {
		logger.Error(errors.New("invalid arguments for post handler"))
		return nil
	}
	handler := Handler{logger: logger, service: service, sessionProvider: sessionProvider, renderer: renderer, commentService: commentService}
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
	data := h.sessionProvider.Get(r, userContextKey)
	user, err := convertionUtil.AnyToGivenType[User](data)
	if err != nil {
		h.logger.Errorf("can not converted session data to user struct with error %s", err.Error())
		return
	}
	ctx := contextUtil.WithContext[User](r.Context(), userContextKey, user)
	_, err = h.service.CreatePost(ctx, CreatePostInput{Content: r.PostFormValue("content")})
	if err != nil {
		h.logger.Error("oops creating post from server")
		h.sessionProvider.Put(r, "create-post-oops", err.Error())
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
		oops.RenderPage(h.renderer, h.logger, h.sessionProvider, r, w, err, http.StatusFound)
		return
	}
	post, err := h.service.GetPostByID(ctx, postID)
	if err != nil {
		h.logger.Error("can not take post from service. Post ID: %s", postID)
		oops.RenderPage(h.renderer, h.logger, h.sessionProvider, r, w, err, http.StatusFound)
		return
	}
	comments, err := h.commentService.GetCommentsByPostID(ctx, postID)
	if err != nil {
		h.logger.Errorf("can not take comments from service. Post ID: %s", postID)
		oops.RenderPage(h.renderer, h.logger, h.sessionProvider, r, w, err, http.StatusFound)
		return
	}
	formattedPost := FormatPost(post)
	h.logger.Debugf("FORMATTED POST: %v", formattedPost)
	formattedComments := FormatComments(comments)
	sessionData := sessionDataFromRequest(h.sessionProvider, r)
	h.Render(w, postData{Post: formattedPost, Session: sessionData, Comments: formattedComments}, http.StatusFound)
}
func (h *Handler) Render(w http.ResponseWriter, data postData, statusCode int) {
	h.logger.Debug("RENDERING THE POST TEMPLATE")
	h.renderer.RenderTemplate(w, h.template, data, statusCode)
}
func (h *Handler) RegisterHandlers(router *mux.Router) {
	router.Handle("/posts", mux.MethodHandler{
		http.MethodPost: h.CreatePost,
	})
	router.Handle("/posts/{postID}", mux.MethodHandler{
		http.MethodGet: h.Show,
	})
}
