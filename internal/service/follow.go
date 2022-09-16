package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	"time"
)

/*FollowUser
- check the followed_id is valid
- take the current user from context
- check if the current user exists ( any login user found)
- if current user.id and followed_id are same return with error
- check if the followed user exists in database by followed_id
- if any error return with error
- if not exist in database return with user not found
- check if there is already a user_follow relation for that users.
- if found
- create user_follow for that users and return created user_follow relationship
*/
func (s *Service) CreateUserFollow(ctx context.Context, followedID string) (*entity.UserFollow, error) {
	if err := validation.IsValidXID(followedID); err != nil {
		s.logger.Error(entity.InvalidUserID)
		return nil, err
	}
	user, exists := contextUtil.FromContext[entity.UserIdentity]("user", ctx)
	if !exists {
		s.logger.Error("can not taken the user from context")
		return nil, entity.UserNotInContext
	}
	if user.ID == followedID {
		s.logger.Debugf("a person can not follow itself.")
		return nil, entity.UserCanNotFollowItself
	}
	exists, err := s.repository.IsUserExistsByID(ctx, followedID)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	if !exists {
		s.logger.Error(entity.UserNotFound)
		return nil, entity.UserNotFound
	}
	exists, err = s.repository.IsUserFollowExists(ctx, user.ID, followedID)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	if exists {
		return nil, entity.AlreadyFollowsTheUser
	}
	createdAt, err := s.repository.CreateUserFollow(ctx, user.ID, followedID)
	if err != nil {
		s.logger.Error("Error creating user follow from repository", s.logger.ErrorModifier(err))
		return nil, err
	}
	message := entity.UserFollowCreatedMessage{
		FollowerID: user.ID,
		FollowedID: followedID,
		CreatedAt:  createdAt,
	}

	err = s.rabbitmqClient.PushMessage(message, "user-follow-created")
	if err != nil {
		s.logger.Error("error pushing user-follow-created message", s.logger.ErrorModifier(err))
		//retry it or sth. because user-follow already created but the counts are not updated....
	}
	return &entity.UserFollow{
		FollowerID: user.ID,
		FollowedID: followedID,
		CreatedAt:  createdAt,
		UpdatedAt:  createdAt,
	}, nil
}
func (s *Service) DeleteUserFollow(ctx context.Context, followedID string) (time.Time, error) {
	if err := validation.IsValidXID(followedID); err != nil {
		s.logger.Error(entity.InvalidUserID)
		return time.Time{}, entity.InvalidUserID
	}
	userIdentity, exists := contextUtil.FromContext[entity.UserIdentity]("user", ctx)
	if !exists {
		s.logger.Error(entity.UserNotInContext)
		return time.Time{}, entity.UserNotInContext
	}

	exists, err := s.repository.IsUserExistsByID(ctx, userIdentity.ID)
	if err != nil {
		s.logger.Error(err)
		return time.Time{}, err
	}
	if !exists {
		s.logger.Error("user can not found in database")
		return time.Time{}, entity.UserNotFound
	}
	exists, err = s.repository.IsUserFollowExists(ctx, userIdentity.ID, followedID)
	if err != nil {
		s.logger.Error(err)
		return time.Time{}, err
	}
	if !exists {
		s.logger.Error(entity.UserFollowNotFound)
		return time.Time{}, entity.UserFollowNotFound
	}
	deletedAt, err := s.repository.DeleteUserFollow(ctx, userIdentity.ID, followedID)
	if err != nil {
		s.logger.Error("can not delete user follow from repository")
		return time.Time{}, err
	}
	message := entity.UserFollowDeletedMessage{
		FollowerID: userIdentity.ID,
		FollowedID: followedID,
		DeletedAt:  deletedAt,
	}
	err = s.rabbitmqClient.PushMessage(message, "user-follow-deleted")
	if err != nil {
		s.logger.Error("error publishing user follow deleted message")
	}
	return deletedAt, nil
}
func (s *Service) IsUserFollowExists(ctx context.Context, followerID, followedID string) (bool, error) {
	return s.repository.IsUserFollowExists(ctx, followerID, followedID)
}
