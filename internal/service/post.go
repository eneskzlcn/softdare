package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/ctxutil"
	"github.com/rs/xid"
	"time"
)

func (s *Service) GetPosts(ctx context.Context, userID string) ([]*entity.Post, error) {
	s.logger.Debugf("GetPosts request for user:", userID)
	return s.repository.GetPosts(ctx, userID)
}

func (s *Service) CreatePost(ctx context.Context, content string) (*entity.Post, error) {
	if err := validation.IsValidContent(content); err != nil {
		s.logger.Error(err)
		return nil, err
	}
	user, exists := ctxutil.FromContext[entity.UserIdentity]("user", ctx)
	if !exists {
		s.logger.Error("unauthorized request user not exist on context")
		return nil, entity.Unauthorized
	}

	id := xid.New().String()

	createdAt, err := s.repository.CreatePost(ctx, id, user.ID, content)
	if err != nil {
		s.logger.Error("oops creating post on repository")
		return nil, err
	}
	message := entity.PostCreatedMessage{
		PostID:    id,
		UserID:    user.ID,
		CreatedAt: createdAt,
	}
	err = s.rabbitmqClient.PushMessage(message, "post-created")
	if err != nil {
		s.logger.Error("error pushing the increase user post count message to the rabbitmq")
		//retry it..
	}
	return &entity.Post{
		ID:           id,
		UserID:       user.ID,
		Content:      content,
		CommentCount: 0,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
		Username:     user.Username,
	}, nil
}

func (s *Service) GetPostByID(ctx context.Context, postID string) (*entity.Post, error) {
	if err := validation.IsValidXID(postID); err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return s.repository.GetPostByID(ctx, postID)
}

func (s *Service) AdjustPostCommentCount(ctx context.Context, postID string, adjustment int) (time.Time, error) {
	if err := validation.IsValidXID(postID); err != nil {
		s.logger.Error("not valid postID : %", postID)
		return time.Time{}, entity.InvalidPostID
	}

	if adjustment <= -10 || adjustment >= 10 {
		s.logger.Error("comment increase amount should be between 1-9 including 1 and 9")
		return time.Time{}, entity.AdjustmentNotValid
	}

	return s.repository.AdjustPostCommentCount(ctx, postID, adjustment)
}

func (s *Service) GetFollowingUsersPosts(ctx context.Context, maxCount int) ([]*entity.Post, error) {
	user, exists := ctxutil.FromContext[entity.UserIdentity]("user", ctx)
	if !exists {
		s.logger.Error(entity.UserNotInContext)
		return nil, entity.UserNotInContext
	}

	followedUserIDs, err := s.repository.GetFollowedUsersOfFollower(ctx, user.ID)
	if err != nil {
		s.logger.Error("error occured when getting followed users of follower")
		return nil, err
	}
	if len(followedUserIDs) == 0 {
		s.logger.Debug("user not following anybody.")
		return s.repository.GetPostsOfGivenUsers(ctx, []string{user.ID}, maxCount)
	}

	usersIncludingCurrentUser := append(followedUserIDs, user.ID)
	followedUsersPosts, err := s.repository.GetPostsOfGivenUsers(ctx, usersIncludingCurrentUser, maxCount)
	if err != nil {
		s.logger.Error("error getting posts of given users from repository", s.logger.ErrorModifier(err))
		return nil, err
	}

	return followedUsersPosts, nil
}
