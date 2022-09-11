package post

import (
	"context"
	"errors"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/oops"
	"github.com/eneskzlcn/softdare/internal/pkg"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
	"github.com/nicolasparada/go-mux"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

type PostService interface {
	CreatePost(ctx context.Context, in CreatePostInput) (*CreatePostResponse, error)
	GetPostByID(ctx context.Context, postID string) (*Post, error)
}
type Renderer interface {
	RenderTemplate(w http.ResponseWriter, template *template.Template, data any, statusCode int)
}
type SessionProvider interface {
	Get(r *http.Request, key string) any
	Put(r *http.Request, key string, data interface{})
	Exists(r *http.Request, key string) bool
}
type Handler struct {
	logger          *zap.SugaredLogger
	service         PostService
	template        *template.Template
	sessionProvider SessionProvider
	renderer        Renderer
}

func NewHandler(logger *zap.SugaredLogger, service PostService, sessionProvider SessionProvider, renderer Renderer) *Handler {
	if logger == nil {
		fmt.Printf("given logger is nil\n")
		return nil
	}
	if service == nil || sessionProvider == nil || renderer == nil {
		logger.Error(errors.New("invalid arguments for post handler"))
		return nil
	}
	handler := Handler{logger: logger, service: service, sessionProvider: sessionProvider, renderer: renderer}
	if err := handler.init(); err != nil {
		logger.Error(err)
		return nil
	}
	return &handler
}
func (h *Handler) init() error {
	tmpl, err := pkg.ParseTemplate("post")
	h.template = tmpl
	if err != nil {
		return err
	}
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
		h.logger.Error("oops creating post from server", zap.Error(err))
		h.sessionProvider.Put(r, "create-post-oops", err.Error())
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func (h *Handler) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := mux.URLParam(ctx, "postID")
	if postID == "" {
		h.logger.Error("can not parse post id", zap.Any("postID", postID))
		return
	}
	post, err := h.service.GetPostByID(ctx, postID)
	if err != nil {
		h.logger.Error("can not take post from service", zap.String("postID", postID))
		oops.RenderPage(h.renderer, h.logger, h.sessionProvider, r, w, err, http.StatusFound)
		return
	}
	formattedPost := FormatPost(post)
	h.Render(w, postData{Post: formattedPost}, http.StatusFound)
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
