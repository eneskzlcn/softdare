package repository

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/lib/pq"
	"time"
)

func (r *Repository) CreatePost(ctx context.Context, postID, userID, content string) (time.Time, error) {
	r.logger.Debugf("A request for creating new post arrived to post repository with id:%s", postID)
	query := `
		INSERT INTO posts (id, user_id, content) 
		VALUES ($1, $2, $3) 
		RETURNING created_at`
	row := r.db.QueryRowContext(ctx, query, postID, userID, content)
	var createdAt time.Time
	err := row.Scan(&createdAt)
	return createdAt, err
}

func (r *Repository) GetPosts(ctx context.Context, userID string) ([]*entity.Post, error) {
	r.logger.Debug("query arrived for all posts")
	query := `
		SELECT posts.id, posts.user_id, posts.content, posts.created_at, posts.updated_at, posts.like_count ,posts.comment_count, users.username
		FROM posts
		INNER JOIN users ON posts.user_id = users.id
		WHERE
    		CASE
        	WHEN $1 <> '' THEN users.id = $1
        	ELSE true
    	END
		ORDER BY posts.id DESC`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*entity.Post
	for rows.Next() {
		var i entity.Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikeCount,
			&i.CommentCount,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Repository) GetPostByID(ctx context.Context, postID string) (*entity.Post, error) {
	r.logger.Debug("Query arrived for get post by id :%", postID)

	query := `
		SELECT posts.id, posts.user_id, posts.content, posts.created_at, posts.updated_at, posts.like_count, posts.comment_count, users.username
		FROM posts 
		INNER JOIN users ON posts.user_id = users.id 
		WHERE posts.id = $1`

	row := r.db.QueryRowContext(ctx, query, postID)
	var i entity.Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LikeCount,
		&i.CommentCount,
		&i.Username,
	)
	if err != nil {
		r.logger.Error("error on getting the post from database")
	}
	return &i, err
}

func (r *Repository) AdjustPostCommentCount(ctx context.Context, postID string, adjustment int) (time.Time, error) {
	query := `
	UPDATE posts
	SET comment_count = comment_count + $1, updated_at = now()
	WHERE id = $2
	RETURNING updated_at;`
	row := r.db.QueryRowContext(ctx, query, adjustment, postID)
	var updatedAt time.Time
	if err := row.Scan(&updatedAt); err != nil {
		r.logger.Error("error scanning query returning")
		return updatedAt, err
	}
	return updatedAt, nil
}

func (r *Repository) GetPostsOfGivenUsers(ctx context.Context, followedUserIDs []string, maxCount int) ([]*entity.Post, error) {
	query := `
		SELECT posts.id, posts.user_id, posts.content, posts.created_at, posts.updated_at, posts.like_count, posts.comment_count, users.username
		FROM posts
		INNER JOIN users ON posts.user_id = users.id
		WHERE posts.user_id = ANY($1)
		ORDER BY posts.created_at DESC
		LIMIT $2
	`
	rows, err := r.db.QueryContext(ctx, query, pq.Array(followedUserIDs), maxCount)
	if err != nil {
		r.logger.Error("can not take posts of given users from database", r.logger.ErrorModifier(err))
		return nil, err
	}
	var items []*entity.Post
	for rows.Next() {
		var i entity.Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikeCount,
			&i.CommentCount,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Repository) AdjustPostLikeCount(ctx context.Context, postID string, adjustment int) (time.Time, error) {
	query := `
		UPDATE posts
		SET like_count = like_count + $1, updated_at = now()
		WHERE id = $2
		RETURNING updated_at;`
	row := r.db.QueryRowContext(ctx, query, adjustment, postID)
	var updatedAt time.Time
	if err := row.Scan(&updatedAt); err != nil {
		r.logger.Error(err)
		return time.Time{}, err
	}
	return updatedAt, nil
}
