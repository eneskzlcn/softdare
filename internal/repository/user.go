package repository

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/entity"
	"time"
)

func (r *Repository) CreateUser(ctx context.Context, userID, email, username string) (time.Time, error) {
	r.logger.Debugf("Creating new user with name %s, id = %s, email = %s", username, userID, email)
	query := `
		INSERT INTO users (id, email, username) 
		VALUES ($1, $2, $3) 
		RETURNING created_at`
	row := r.db.QueryRowContext(ctx, query, userID, email, username)
	var createdAt time.Time
	err := row.Scan(&createdAt)
	return createdAt, err
}

func (r *Repository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	query := `
		SELECT id, email, username, post_count, follower_count, followed_count, created_at, updated_at 
		FROM users 
		WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, userID)
	var i entity.User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.PostCount,
		&i.FollowerCount,
		&i.FollowedCount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `
		SELECT id, email, username, post_count, follower_count, followed_count, created_at, updated_at 
		FROM users 
		WHERE username = $1`
	row := r.db.QueryRowContext(ctx, query, username)
	var i entity.User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.PostCount,
		&i.FollowerCount,
		&i.FollowedCount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, email, username, post_count, follower_count, followed_count ,created_at, updated_at 
		FROM users 
		WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)
	var i entity.User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.PostCount,
		&i.FollowerCount,
		&i.FollowedCount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

func (r *Repository) IsUserExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := `
	SELECT EXISTS ( 
		SELECT 1 
		FROM users 
		WHERE email = $1) `
	row := r.db.QueryRowContext(ctx, query, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

func (r *Repository) IsUserExistsByID(ctx context.Context, userID string) (bool, error) {
	query := `
	SELECT EXISTS ( 
		SELECT 1 
		FROM users 
		WHERE id = $1) `
	row := r.db.QueryRowContext(ctx, query, userID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

func (r *Repository) IsUserExistsByUsername(ctx context.Context, username string) (bool, error) {
	query := `
	SELECT EXISTS (
		SELECT 1 
		FROM users 
		WHERE username ILIKE $1 )`
	row := r.db.QueryRowContext(ctx, query, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

func (r *Repository) AdjustUserPostCount(ctx context.Context, userID string, adjustment int) (time.Time, error) {
	query := `
		UPDATE users
		SET post_count = post_count + $1, updated_at = now()
		WHERE id = $2
		RETURNING updated_at;`
	row := r.db.QueryRowContext(ctx, query, adjustment, userID)
	var updatedAt time.Time
	err := row.Scan(&updatedAt)
	return updatedAt, err
}

func (r *Repository) AdjustUserFollowerCount(ctx context.Context, userID string, adjustment int) (time.Time, error) {
	query := `
	UPDATE users
	SET follower_count = follower_count + $1, updated_at = now()
	WHERE id = $2
	RETURNING updated_at;`
	row := r.db.QueryRowContext(ctx, query, adjustment, userID)
	var updatedAt time.Time
	err := row.Scan(&updatedAt)
	return updatedAt, err
}

func (r *Repository) AdjustUserFollowedCount(ctx context.Context, userID string, adjustment int) (time.Time, error) {
	query := `
	UPDATE users
	SET followed_count = followed_count + $1, updated_at = now()
	WHERE id = $2
	RETURNING updated_at;`
	row := r.db.QueryRowContext(ctx, query, adjustment, userID)
	var updatedAt time.Time
	err := row.Scan(&updatedAt)
	return updatedAt, err
}

func (r *Repository) DeleteUserFollow(ctx context.Context, followerID, followedID string) (time.Time, error) {
	query := `
	DELETE from user_follows
	WHERE follower_id = $1 AND followed_id = $2
	RETURNING now()::timestamp AS deleted_at;
	`
	row := r.db.QueryRowContext(ctx, query, followerID, followedID)
	var deletedAt time.Time
	if err := row.Scan(&deletedAt); err != nil {
		r.logger.Error(err)
		return time.Time{}, err
	}
	return deletedAt, nil
}

func (r *Repository) GetUsersWithFollowedOptionHasUsernameSimilarTo(ctx context.Context, usernameSimilarTo string, userID string) ([]*entity.UserWithFollowedOption, error) {
	query := `
	SELECT id, email, username, post_count, follower_count, followed_count, created_at, updated_at,
		(SELECT EXISTS ( SELECT 1 FROM user_follows uf WHERE uf.follower_id = $1 and uf.followed_id = u.id )) as isFollowed
	FROM users u
	WHERE u.username similar to $2`

	sqlParameterFormOfUsernameSimilarTo := "%" + usernameSimilarTo + "%"

	rows, err := r.db.QueryContext(ctx, query, userID, sqlParameterFormOfUsernameSimilarTo)
	defer rows.Close()
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}

	usersWithFollowedOption := make([]*entity.UserWithFollowedOption, 0)
	for rows.Next() {
		var user entity.UserWithFollowedOption
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Username,
			&user.PostCount,
			&user.FollowerCount,
			&user.FollowedCount,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.IsFollowed)
		if err != nil {
			r.logger.Error(err)
			return nil, err
		}
		usersWithFollowedOption = append(usersWithFollowedOption, &user)
	}
	if err = rows.Err(); err != nil {
		r.logger.Error(err)
		return nil, err
	}
	if err = rows.Close(); err != nil {
		r.logger.Error(err)
		return nil, err
	}
	return usersWithFollowedOption, nil
}
