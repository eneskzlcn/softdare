package web

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	"net/http"
)

type profileData struct {
	IsFollowedUser bool
	Session        CommonSessionData
	User           entity.User
	Posts          []entity.FormattedPost
}

func (h *Handler) ShowProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := h.urlParamExtractor(ctx, "username")
	if err := validation.IsValidUsername(username); err != nil {
		h.logger.Error("Invalid username for profile", h.logger.ErrorModifier(err))
		return
	}
	user, err := h.service.GetUserByUsername(ctx, username)
	if err != nil {
		h.logger.Error("error getting user from service", h.logger.StringModifier("username", username), h.logger.ErrorModifier(err))
		return
	}
	posts, err := h.service.GetFormattedPosts(ctx, user.ID)
	if err != nil {
		h.logger.Error("error getting posts from service", h.logger.StringModifier("userID", user.ID), h.logger.ErrorModifier(err))
		return
	}
	isLoggedIn, userIdentity := h.CommonSessionDataFromRequest(r)
	isFollowedUser := false
	if isLoggedIn && userIdentity.ID != user.ID {
		exists, err := h.service.IsUserFollowExists(ctx, userIdentity.ID, user.ID)
		if err != nil {
			h.logger.Error("error checking is user follow exists")
		}
		if exists {
			isFollowedUser = true
		}
	}
	h.RenderProfile(w, profileData{User: *user, Posts: posts, IsFollowedUser: isFollowedUser, Session: CommonSessionData{IsLoggedIn: isLoggedIn, User: userIdentity}}, http.StatusFound)
}
func (h *Handler) CreateUserFollow(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.logger.Error("can not parse request form", h.logger.ErrorModifier(err))
		return
	}
	followedUserID := r.PostFormValue("userID")
	if err := validation.IsValidXID(followedUserID); err != nil {
		h.logger.Error("invalid user id to follow")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	isLoggedIn, userIdentity := h.CommonSessionDataFromRequest(r)
	if !isLoggedIn {
		h.logger.Error("not logged in user can not follow anybody")
		http.Redirect(w, r, "/login", http.StatusBadRequest)
		return
	}
	ctx := context.WithValue(r.Context(), "user", userIdentity)
	_, err := h.service.FollowUser(ctx, followedUserID)
	if err != nil {
		h.logger.Error("error creating user follow from service layer", h.logger.ErrorModifier(err))
		http.Redirect(w, r, "/", http.StatusFound)
	}
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
func (h *Handler) RenderProfile(w http.ResponseWriter, data profileData, statusCode int) {
	h.RenderPage("profile", w, data, statusCode)
}
