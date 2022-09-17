package web

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/convertutil"
	"github.com/rs/xid"
	"net/http"
)

func (h *Handler) CreateComment(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		h.handleCreateCommentError(w, req, err)
		return
	}
	postID := req.PostFormValue("post_id")
	content := req.PostFormValue("content")
	data := h.session.Get(req, "user")
	user, err := convertutil.AnyTo[entity.UserIdentity](data)
	if err != nil {
		h.logger.Error("error getting user from session")
		h.handleCreateCommentError(w, req, err)
		return
	}
	_, err = xid.FromString(user.ID)
	if err != nil {
		h.logger.Error("invalid user id for create comment")
		h.handleCreateCommentError(w, req, err)
		return
	}
	ctx := context.WithValue(req.Context(), "user", user)
	_, err = h.service.CreateComment(ctx, postID, content)
	if err != nil {
		h.handleCreateCommentError(w, req, err)
		return
	}

	http.Redirect(w, req, req.Referer(), http.StatusFound)
}

func (h *Handler) handleCreateCommentError(w http.ResponseWriter, req *http.Request, err error) {
	h.logger.Error("error creating comment on service")
	h.session.Put(req, "create-comment-error", err.Error())
	h.session.Put(req, "create-comment-form", req.PostForm)
	http.Redirect(w, req, req.Referer(), http.StatusFound)
}
