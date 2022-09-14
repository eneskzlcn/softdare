package web

import (
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	"net/http"
)

type profileData struct {
	Session CommonSessionData
	User    entity.User
	Posts   []entity.FormattedPost
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

	h.RenderProfile(w, profileData{User: *user, Posts: posts, Session: CommonSessionData{IsLoggedIn: isLoggedIn, User: userIdentity}}, http.StatusFound)
}

func (h *Handler) RenderProfile(w http.ResponseWriter, data profileData, statusCode int) {
	h.RenderPage("profile", w, data, statusCode)
}
