package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/trend"
	"github.com/eneskzlcn/softdare/internal/entity"
	customerror "github.com/eneskzlcn/softdare/internal/error"
)

/*GetTrendPostsByGivenTrendFindingStrategy function works like
- Looks the cache data first, if found returns the cached data
- Search from database by given finding strategy.
- Save the results to the cache
- Then return the data.

*/
func (s *Service) GetTrendPostsByGivenTrendFindingStrategy(ctx context.Context, strategy trend.FindingStrategy) ([]*entity.Post, error) {
	switch strategy {
	case trend.LikeCountFindingStrategy:
		return s.GetTrendPostsByLikeCountFindingStrategy(ctx)
	default:
		return nil, customerror.NotExistingTrendFoundStrategy
	}
	return nil, customerror.NotExistingTrendFoundStrategy
}
func (s *Service) GetTrendPostsByLikeCountFindingStrategy(ctx context.Context) ([]*entity.Post, error) {
	return s.repository.GetPostsHasLikeCountMoreThanGiven(ctx, trend.MinimumLikeCountToAcceptAsTrendForLikeCountFindingStrategy)
}
