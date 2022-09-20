package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/search"
	"github.com/eneskzlcn/softdare/internal/entity"
	customerror "github.com/eneskzlcn/softdare/internal/error"
	"github.com/eneskzlcn/softdare/internal/util/ctxutil"
	"github.com/eneskzlcn/softdare/internal/util/timeutil"
	"github.com/eneskzlcn/softdare/internal/util/userutil"
)

func (s *Service) SearchUserByGivenSearchCriteria(ctx context.Context, searchContent string, criteria search.Criteria) ([]entity.FormattedUserWithFollowedOption, error) {
	switch criteria {
	case search.UserByUsernameCriteria:
		return s.SearchUserByUsernameSearchCriteria(ctx, searchContent)
		break
	default:
		s.logger.Error("not valid search criteria")
		return nil, nil
	}
	return nil, nil
}
func (s *Service) SearchUserByUsernameSearchCriteria(ctx context.Context, searchContent string) ([]entity.FormattedUserWithFollowedOption, error) {
	userIdentity, exists := ctxutil.FromContext[entity.UserIdentity]("user", ctx)
	if !exists {
		s.logger.Error(customerror.UserNotInContext)
		return nil, customerror.UserNotInContext
	}
	searchedUsersWithFollowedOption, err := s.repository.GetUsersWithFollowedOptionHasUsernameSimilarTo(ctx, searchContent, userIdentity.ID)
	if err != nil {
		s.logger.Error(customerror.SearchUserWithFollowingOptionByUsername)
		return nil, customerror.SearchUserWithFollowingOptionByUsername
	}
	formattedSearchedUsersWithFollowedOptions := userutil.FormatUsersWithFollowingOption(searchedUsersWithFollowedOption, timeutil.ToAgoFormatter)
	return formattedSearchedUsersWithFollowedOptions, nil
}
