package web

import (
	"github.com/eneskzlcn/softdare/internal/core/trend"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/postutil"
	"github.com/eneskzlcn/softdare/internal/util/timeutil"
	"net/http"
)

type trendPageData struct {
	Session trendsSessionData
	Posts   []entity.FormattedPost
}
type trendsSessionData struct {
	IsLoggedIn bool
	User       entity.UserIdentity
}

func (h *Handler) ShowTrends(w http.ResponseWriter, r *http.Request) {
	trendSessionData := h.GetTrendsSessionData(r)
	posts, err := h.service.GetTrendPostsByGivenTrendFindingStrategy(r.Context(), trend.LikeCountFindingStrategy)
	if err != nil {
		h.logger.Error(err)
		return
	}
	formattedPosts := postutil.FormatPosts(posts, timeutil.ToAgoFormatter)
	h.RenderTrends(w, trendPageData{Session: trendSessionData, Posts: formattedPosts}, http.StatusFound)
}

func (h *Handler) RenderTrends(w http.ResponseWriter, data trendPageData, statusCode int) {
	h.RenderPage("trends", w, data, statusCode)
}
func (h *Handler) GetTrendsSessionData(r *http.Request) trendsSessionData {
	isLoggedIn, userIdentity := h.CommonSessionDataFromRequest(r)
	return trendsSessionData{
		IsLoggedIn: isLoggedIn,
		User:       userIdentity,
	}
}
