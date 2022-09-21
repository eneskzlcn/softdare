package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/trend"
	"github.com/eneskzlcn/softdare/internal/entity"
	customerror "github.com/eneskzlcn/softdare/internal/error"
	"github.com/eneskzlcn/softdare/internal/util/convertutil"
	"github.com/eneskzlcn/softdare/internal/util/postutil"
	"time"
)

const (
	TrendPostsCacheKey            = "trend-posts"
	TrendPostsCacheExpireDuration = time.Second * 60
)

/*GetTrendPostsByGivenTrendFindingStrategy function works like
- Looks the cache data first, if found returns the cached data
- Search from database by given finding strategy.
- Save the results to the cache
- Then return the data.
*/
func (s *Service) GetTrendPostsByGivenTrendFindingStrategy(ctx context.Context, strategy trend.FindingStrategy) ([]*entity.Post, error) {
	postsAny, err := s.cache.Get(TrendPostsCacheKey)
	if err == nil {
		s.logger.Debug("trend posts found from cache...")
		posts, err := convertutil.AnyTo[[]entity.Post](postsAny)
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}
		return postutil.ToPostsPtr(posts), nil
	}
	s.logger.Debug("trend posts can not found from cache...")
	switch strategy {
	case trend.LikeCountFindingStrategy:
		return s.GetTrendPostsByLikeCountFindingStrategy(ctx)
	default:
		return nil, customerror.NotExistingTrendFoundStrategy
	}
}
func (s *Service) GetTrendPostsByLikeCountFindingStrategy(ctx context.Context) ([]*entity.Post, error) {
	posts, err := s.repository.GetPostsHasLikeCountMoreThanGiven(ctx,
		trend.MinimumLikeCountToAcceptAsTrendForLikeCountFindingStrategy)
	if err != nil {
		return nil, err
	}
	err = s.cache.SetWithExpire(TrendPostsCacheKey, postutil.ToPostValue(posts), TrendPostsCacheExpireDuration)
	if err != nil {
		s.logger.Debug("error occurred when caching the trending posts", err)
	}
	return posts, nil
}
