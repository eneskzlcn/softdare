package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/ctxutil"
	"github.com/rs/xid"
	"time"
)

func (s *Service) CreateComment(ctx context.Context, postID, content string) (*entity.Comment, error) {
	if err := validation.IsValidXID(postID); err != nil {
		s.logger.Error(err)
		return nil, err
	}
	if err := validation.IsValidContent(content); err != nil {
		s.logger.Error(err)
		return nil, err
	}
	user, exists := ctxutil.FromContext[entity.UserIdentity]("user", ctx)
	if !exists {
		s.logger.Error("%s , Exists: %t", entity.CouldNotTakeUserFromContext, exists)
		return nil, entity.CouldNotTakeUserFromContext
	}
	id := xid.New().String()
	createdAt, err := s.repository.CreateComment(ctx, id, user.ID, postID, content)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	err = s.rabbitmqClient.PushMessage(entity.CommentCreatedMessage{
		CommentID: id,
		PostID:    postID,
		CreatedAt: createdAt,
	}, "comment-created")
	if err != nil {
		s.logger.Error("error publishing increase-post-comment-count message")
	}
	return &entity.Comment{
		ID:        id,
		PostID:    postID,
		UserID:    user.ID,
		Content:   content,
		LikeCount: 0,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Username:  user.Username,
	}, nil
}

func (s *Service) GetCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error) {
	if err := validation.IsValidXID(postID); err != nil {
		s.logger.Error(err)
		return nil, err
	}
	comments, err := s.repository.GetCommentsByPostID(ctx, postID)
	if err != nil {
		s.logger.Errorf("error getting comments from repository. Error: %s", err.Error())
		return nil, err
	}
	return comments, nil
}

func (s *Service) AdjustCommentLikeCount(ctx context.Context, commentID string, adjustment int) (time.Time, error) {
	if err := validation.IsValidXID(commentID); err != nil {
		s.logger.Error(err)
		return time.Time{}, err
	}
	if adjustment <= CommentLikeAdjustmentLowerBound || adjustment >= CommentLikeAdjustmentUpperBound {
		s.logger.Error(entity.AdjustmentNotValid)
		return time.Time{}, entity.AdjustmentNotValid
	}
	return s.repository.AdjustCommentLikeCount(ctx, commentID, adjustment)
}
