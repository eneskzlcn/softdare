package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	"github.com/rs/xid"
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
	user, exists := contextUtil.FromContext[entity.UserIdentity]("user", ctx)
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
