package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/ctxutil"
	"time"
)

func (s *Service) CreatePostLike(ctx context.Context, postID string) (time.Time, error) {
	if err := validation.IsValidXID(postID); err != nil {
		s.logger.Error(err)
		return time.Time{}, err
	}

	userIdentity, exists := ctxutil.FromContext[entity.UserIdentity]("user", ctx)
	if !exists {
		s.logger.Error(entity.CouldNotTakeUserFromContext)
		return time.Time{}, entity.CouldNotTakeUserFromContext
	}

	createdAt, err := s.repository.CreatePostLike(ctx, userIdentity.ID, postID)
	if err != nil {
		s.logger.Error(err)
		return time.Time{}, err
	}
	//publish a post-like-created message.
	return createdAt, nil
}
