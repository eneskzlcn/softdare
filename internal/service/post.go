package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	postUtil "github.com/eneskzlcn/softdare/internal/util/post"
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
	user, exists := contextUtil.FromContext[entity.UserIdentity]("user", ctx)
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
	increaseUserPostCountMessage := entity.PostCreatedMessage{
		PostID: id,
		UserID: user.ID,
	}
	err = s.rabbitmqClient.PushMessage(increaseUserPostCountMessage, "post-created")
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
func (s *Service) IncreasePostCommentCount(ctx context.Context, postID string, increaseAmount int) (time.Time, error) {
	if err := validation.IsValidXID(postID); err != nil {
		s.logger.Error("not valid postID : %", postID)
		return time.Time{}, err
	}
	if increaseAmount <= 0 || increaseAmount >= 10 {
		s.logger.Error("comment increase amount should be between 1-9 including 1 and 9")
		return time.Time{}, entity.IncreaseAmountNotValid
	}
	return s.repository.IncreasePostCommentCount(ctx, postID, increaseAmount)
}
func (s *Service) GetFormattedPosts(ctx context.Context, userID string) ([]entity.FormattedPost, error) {
	posts, err := s.GetPosts(ctx, userID)
	if err != nil {
		s.logger.Error("oops getting posts from post service")
		return nil, err
	}
	formattedPosts := make([]entity.FormattedPost, 0)
	for _, postPtr := range posts {
		formattedPost := entity.FormattedPost{
			ID:           postPtr.ID,
			Username:     postPtr.Username,
			Content:      postPtr.Content,
			CommentCount: postPtr.CommentCount,
		}
		formattedPost.CreatedAt = postUtil.FormatPostTime(postPtr.CreatedAt)
		formattedPosts = append(formattedPosts, formattedPost)
	}
	return formattedPosts, nil
}
