package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	customerror "github.com/eneskzlcn/softdare/internal/error"
	"github.com/eneskzlcn/softdare/internal/message"
	"github.com/eneskzlcn/softdare/internal/util/ctxutil"
	"time"
)

const (
	PostLikeAdjustmentLowerBound    = -10
	PostLikeAdjustmentUpperBound    = 10
	CommentLikeAdjustmentLowerBound = -10
	CommentLikeAdjustmentUpperBound = 10
)

func (s *Service) CreatePostLike(ctx context.Context, postID string) (time.Time, error) {
	if err := validation.IsValidXID(postID); err != nil {
		s.logger.Error(err)
		return time.Time{}, err
	}

	userIdentity, exists := ctxutil.FromContext[entity.UserIdentity]("user", ctx)
	if !exists {
		s.logger.Error(customerror.CouldNotTakeUserFromContext)
		return time.Time{}, customerror.CouldNotTakeUserFromContext
	}
	exists, err := s.repository.IsPostLikeExists(ctx, postID, userIdentity.ID)
	if err != nil {
		s.logger.Error(err)
		return time.Time{}, err
	}
	if exists {
		s.logger.Error(customerror.UserAlreadyLikedThePost)
		return time.Time{}, customerror.UserAlreadyLikedThePost
	}
	createdAt, err := s.repository.CreatePostLike(ctx, userIdentity.ID, postID)
	if err != nil {
		s.logger.Error(err)
		return time.Time{}, err
	}
	adjustPostLikeCountMessage := message.PostLikeCreated{
		PostID:    postID,
		UserID:    userIdentity.ID,
		CreatedAt: createdAt,
	}
	err = s.rabbitmqClient.PushMessage(adjustPostLikeCountMessage, "post-like-created")
	if err != nil {
		s.logger.Error(err)
	}
	return createdAt, nil
}

func (s *Service) CreateCommentLike(ctx context.Context, commentID string) (time.Time, error) {
	if err := validation.IsValidXID(commentID); err != nil {
		s.logger.Error(err)
		return time.Time{}, err
	}
	userIdentity, exists := ctxutil.FromContext[entity.UserIdentity]("user", ctx)
	if !exists {
		s.logger.Error(customerror.CouldNotTakeUserFromContext)
		return time.Time{}, customerror.CouldNotTakeUserFromContext
	}
	exists, err := s.repository.IsCommentLikeExists(ctx, commentID, userIdentity.ID)
	if err != nil {
		s.logger.Error(err)
		return time.Time{}, err
	}
	if exists {
		s.logger.Error(customerror.UserAlreadyLikedTheComment)
		return time.Time{}, customerror.UserAlreadyLikedTheComment
	}

	createdAt, err := s.repository.CreateCommentLike(ctx, commentID, userIdentity.ID)
	if err != nil {
		s.logger.Error(err)
		return time.Time{}, err
	}

	commentLikeCreatedMessage := message.CommentLikeCreated{
		CommentID: commentID,
		UserID:    userIdentity.ID,
		CreatedAt: createdAt,
	}
	err = s.rabbitmqClient.PushMessage(commentLikeCreatedMessage, "comment-like-created")
	if err != nil {
		s.logger.Error("Can not push comment like created message.", err)
	}
	return createdAt, nil
}
