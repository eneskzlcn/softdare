package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
)

func (s *Service) FollowUser(ctx context.Context, followerID, followedID string) (*entity.UserFollow, error) {
	if err := validation.IsValidXID(followerID); err != nil {
		s.logger.Error(entity.InvalidUserID)
		return nil, err
	}
	if err := validation.IsValidXID(followedID); err != nil {
		s.logger.Error(entity.InvalidUserID)
		return nil, err
	}
	createdAt, err := s.repository.FollowUser(ctx, followerID, followedID)
	if err != nil {
		s.logger.Error("Error creating user follow from repository", s.logger.ErrorModifier(err))
		return nil, err
	}
	return &entity.UserFollow{
		FollowerID: followerID,
		FollowedID: followedID,
		CreatedAt:  createdAt,
		UpdatedAt:  createdAt,
	}, nil
}
