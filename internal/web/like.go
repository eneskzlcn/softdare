package web

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/entity"
	"net/http"
)

func (h *Handler) CreatePostLike(w http.ResponseWriter, r *http.Request) {
	isLoggedIn, userSessionData := h.CommonSessionDataFromRequest(r)
	if !isLoggedIn {
		h.logger.Error(entity.NotLoggedInUser)
		h.session.Put(r, "come-from-home", true)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		h.logger.Error(err)
		h.ShowOops(w, r, err, http.StatusFound)
		return
	}
	postID := r.PostFormValue("postID")

	ctx := context.WithValue(r.Context(), "user", userSessionData)
	_, err := h.service.CreatePostLike(ctx, postID)
	if err != nil {
		h.logger.Error(err)
		h.ShowOops(w, r, err, http.StatusFound)
		return
	}
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}

func (h *Handler) CreateCommentLike(w http.ResponseWriter, r *http.Request) {
	isLoggedIn, userSessionData := h.CommonSessionDataFromRequest(r)
	if !isLoggedIn {
		h.logger.Error(entity.NotLoggedInUser)
		h.session.Put(r, "come-from-home", true)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		h.logger.Error(err)
		h.ShowOops(w, r, err, http.StatusFound)
		return
	}
	commentID := r.PostFormValue("commentID")

	ctx := context.WithValue(r.Context(), "user", userSessionData)

	_, err := h.service.CreateCommentLike(ctx, commentID)
	if err != nil {
		h.logger.Error(err)
		h.ShowOops(w, r, err, http.StatusFound)
		return
	}
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
