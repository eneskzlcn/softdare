package web

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/convertion"
	"github.com/nicolasparada/go-mux"
	"net/http"
	"net/url"
)

type postData struct {
	Session  postSessionData
	Post     entity.FormattedPost
	Comments []entity.FormattedComment
}
type postSessionData struct {
	IsLoggedIn         bool
	User               entity.UserIdentity
	CreateCommentForm  url.Values
	CreateCommentError error
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("CREATE POST REQUEST ARRIVED")
	if err := r.ParseForm(); err != nil {
		h.logger.Error(err)
		return
	}
	data := h.session.Get(r, "user")
	user, err := convertion.AnyToGivenType[entity.UserIdentity](data)
	if err != nil {
		h.logger.Errorf("can not converted session data to user struct with error %s", err.Error())
		return
	}
	ctx := context.WithValue(r.Context(), "user", user)
	_, err = h.service.CreatePost(ctx, r.PostFormValue("content"))
	if err != nil {
		h.logger.Error("oops creating post from server")
		h.session.Put(r, "create-post-oops", err.Error())
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func (h *Handler) ShowPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID := mux.URLParam(ctx, "postID")
	if err := validation.IsValidXID(postID); err != nil {
		h.ShowOops(w, r, err, http.StatusFound)
		return
	}
	post, err := h.service.GetPostByID(ctx, postID)
	if err != nil {
		h.ShowOops(w, r, err, http.StatusFound)
		return
	}
	comments, err := h.service.GetCommentsByPostID(ctx, postID)
	if err != nil {
		h.ShowOops(w, r, err, http.StatusFound)
		return
	}
	formattedPost := entity.FormatPost(post)
	formattedComments := entity.FormatComments(comments)
	sessionData := h.GetPostSessionData(r)
	h.RenderPost(w, postData{Post: formattedPost, Session: sessionData, Comments: formattedComments}, http.StatusFound)
}
func (h *Handler) RenderPost(w http.ResponseWriter, data postData, status int) {
	h.RenderPage("post", w, data, status)
}
func (h *Handler) GetPostSessionData(r *http.Request) (out postSessionData) {
	isLoggedIn, user := h.CommonSessionDataFromRequest(r)
	if isLoggedIn {
		out.IsLoggedIn = isLoggedIn
		out.User = user
	}
	if h.session.Exists(r, "create-comment-error") {
		out.CreateCommentError = h.session.PopError(r, "create-comment-error")
	}
	if h.session.Exists(r, "create-comment-form") {
		form := h.session.Get(r, "create-comment-form")
		urlForm, err := convertion.AnyToGivenType[url.Values](form)
		if err == nil {
			out.CreateCommentForm = urlForm
		}
	}
	return
}
