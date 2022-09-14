package web

import (
	"context"
	"encoding/gob"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/router"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/entity"
	"html/template"
	"net/http"
	"net/url"
	"time"
)

type PostService interface {
	GetPosts(ctx context.Context, userID string) ([]*entity.Post, error)
	GetFormattedPosts(ctx context.Context, userID string) ([]entity.FormattedPost, error)
	CreatePost(ctx context.Context, content string) (*entity.Post, error)
	GetPostByID(ctx context.Context, postID string) (*entity.Post, error)
	IncreasePostCommentCount(ctx context.Context, postID string, increaseAmount int) (time.Time, error)
}
type CommentService interface {
	CreateComment(ctx context.Context, postID, content string) (*entity.Comment, error)
	GetCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error)
}
type LoginService interface {
	Login(ctx context.Context, email string, username *string) (*entity.User, error)
}
type ProfileService interface {
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}
type Service interface {
	PostService
	CommentService
	LoginService
	ProfileService
}
type Renderer interface {
	Render(w http.ResponseWriter, tmpl *template.Template, data any, statusCode int)
}
type PageTemplates map[string]*template.Template

type Handler struct {
	logger            logger.Logger
	session           session.Session
	handler           http.Handler
	service           Service
	renderer          Renderer
	templates         PageTemplates
	urlParamExtractor func(ctx context.Context, key string) string
}

func NewHandler(logger logger.Logger, session session.Session, service Service, renderer Renderer) *Handler {
	if logger == nil {
		return nil
	}
	if session == nil || service == nil || renderer == nil {
		logger.Error(entity.InvalidConstructorArguments)
		return nil
	}
	handler := Handler{logger: logger, session: session, service: service, renderer: renderer}

	handler.init()

	return &handler
}
func (h *Handler) init() {
	muxRouter := router.NewMuxRouterAdapter()
	h.urlParamExtractor = muxRouter.ExtractURLParam
	h.handler = muxRouter
	h.handler = h.session.Enable(h.handler)
	h.RegisterHandlers(muxRouter)

	gob.Register(url.Values{})
	gob.Register(entity.UserIdentity{})

	templates := make(PageTemplates, 0)
	templates["home"] = ParseTemplate("home.gohtml")
	templates["login"] = ParseTemplate("login.gohtml")
	templates["post"] = ParseTemplate("post.gohtml")
	templates["oops"] = ParseTemplate("oops.gohtml")
	templates["profile"] = ParseTemplate("profile.gohtml")
	h.templates = templates
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.handler.ServeHTTP(w, req)
}
func (h *Handler) RenderPage(page string, w http.ResponseWriter, data any, statusCode int) {
	h.renderer.Render(w, h.templates[page], data, statusCode)
}
func (h *Handler) RegisterHandlers(muxRouter router.Router) {
	muxRouter.Handle("/", router.MethodHandlers{
		http.MethodGet: h.ShowHome,
	})
	muxRouter.Handle("/comments", router.MethodHandlers{
		http.MethodPost: h.CreateComment,
	})
	muxRouter.Handle("/login", router.MethodHandlers{
		http.MethodGet:  h.ShowLogin,
		http.MethodPost: h.Login,
	})
	muxRouter.Handle("/logout", router.MethodHandlers{
		http.MethodPost: h.Logout,
	})
	muxRouter.Handle("/posts", router.MethodHandlers{
		http.MethodPost: h.CreatePost,
	})
	muxRouter.Handle("/posts/{postID}", router.MethodHandlers{
		http.MethodGet: h.ShowPost,
	})
	muxRouter.Handle("/{username}", router.MethodHandlers{
		http.MethodGet: h.ShowProfile,
	})
}
