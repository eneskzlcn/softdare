package web

import (
	"context"
	"encoding/gob"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/middleware"
	"github.com/eneskzlcn/softdare/internal/core/router"
	"github.com/eneskzlcn/softdare/internal/core/search"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/core/trend"
	"github.com/eneskzlcn/softdare/internal/entity"
	customerror "github.com/eneskzlcn/softdare/internal/error"
	"html/template"
	"net/http"
	"net/url"
	"time"
)

type PostService interface {
	GetPosts(ctx context.Context, userID string) ([]*entity.Post, error)
	CreatePost(ctx context.Context, content string) (*entity.Post, error)
	GetPostByID(ctx context.Context, postID string) (*entity.Post, error)
	AdjustPostCommentCount(ctx context.Context, postID string, increaseAmount int) (time.Time, error)
	GetFollowingUsersPosts(ctx context.Context, maxCount int) ([]*entity.Post, error)
}

type UserService interface {
	Login(ctx context.Context, email string, username *string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	SearchUserByGivenSearchCriteria(ctx context.Context, searchContent string, criteria search.Criteria) ([]entity.FormattedUserWithFollowedOption, error)
}

type CommentService interface {
	CreateComment(ctx context.Context, postID, content string) (*entity.Comment, error)
	GetCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error)
}

type FollowService interface {
	CreateUserFollow(ctx context.Context, followedID string) (*entity.UserFollow, error)
	IsUserFollowExists(ctx context.Context, followerID, followedID string) (bool, error)
	DeleteUserFollow(ctx context.Context, followedID string) (time.Time, error)
}

type LikeService interface {
	CreatePostLike(ctx context.Context, postID string) (time.Time, error)
	CreateCommentLike(ctx context.Context, commentID string) (time.Time, error)
}
type TrendService interface {
	GetTrendPostsByGivenTrendFindingStrategy(ctx context.Context, strategy trend.FindingStrategy) ([]*entity.Post, error)
}
type Service interface {
	PostService
	CommentService
	UserService
	FollowService
	LikeService
	TrendService
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
		logger.Error(customerror.InvalidConstructorArguments)
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
	h.RegisterHandlers(muxRouter)

	h.applyMiddlewares()

	gob.Register(url.Values{})
	gob.Register(entity.UserIdentity{})

	templates := make(PageTemplates, 0)
	templates["home"] = ParseTemplate("home.gohtml")
	templates["login"] = ParseTemplate("login.gohtml")
	templates["post"] = ParseTemplate("post.gohtml")
	templates["oops"] = ParseTemplate("oops.gohtml")
	templates["profile"] = ParseTemplate("profile.gohtml")
	templates["search"] = ParseTemplate("search.gohtml")
	templates["trends"] = ParseTemplate("trends.gohtml")
	h.templates = templates
}

func (h *Handler) applyMiddlewares() {
	//apply session middleware
	h.handler = h.session.Enable(h.handler)
	//apply method overriding middleware
	h.handler = middleware.OverrideFormMethods(h.handler)
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

	muxRouter.Handle("/posts/@{postID}", router.MethodHandlers{
		http.MethodGet: h.ShowPost,
	})

	muxRouter.Handle("/@{username}", router.MethodHandlers{
		http.MethodGet: h.ShowProfile,
	})

	muxRouter.Handle("/follow", router.MethodHandlers{
		http.MethodPost: h.CreateUserFollow,
	})

	muxRouter.Handle("/unfollow", router.MethodHandlers{
		http.MethodDelete: h.DeleteUserFollow,
	})

	muxRouter.Handle("/like/post", router.MethodHandlers{
		http.MethodPost: h.CreatePostLike,
	})

	muxRouter.Handle("/like/comment", router.MethodHandlers{
		http.MethodPost: h.CreateCommentLike,
	})

	muxRouter.Handle("/search", router.MethodHandlers{
		http.MethodPost: h.Search,
	})

	muxRouter.Handle("/trends", router.MethodHandlers{
		http.MethodGet: h.ShowTrends,
	})
}
