package web

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/search"
	"github.com/eneskzlcn/softdare/internal/entity"
	"net/http"
)

type searchPageData struct {
	Session CommonSessionData
	Users   []entity.FormattedUserWithFollowedOption
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.logger.Error(err)
		return
	}
	searchData := r.PostFormValue("searchData")

	searchSessionData := h.GetSearchSessionDataFromRequest(r)

	ctx := context.WithValue(r.Context(), "user", searchSessionData.User)
	formattedSearchedUsersWithIsFollowedOption, err :=
		h.service.SearchUserByGivenSearchCriteria(ctx, searchData, search.UserByUsernameCriteria)
	if err != nil {
		h.logger.Error(err)
		return
	}
	h.RenderSearch(w, searchPageData{Session: searchSessionData, Users: formattedSearchedUsersWithIsFollowedOption}, http.StatusFound)
}

func (h *Handler) RenderSearch(w http.ResponseWriter, data searchPageData, statusCode int) {
	h.RenderPage("search", w, data, statusCode)
}
func (h *Handler) GetSearchSessionDataFromRequest(r *http.Request) CommonSessionData {
	isLoggedIn, userIdentity := h.CommonSessionDataFromRequest(r)
	return CommonSessionData{IsLoggedIn: isLoggedIn, User: userIdentity}
}
