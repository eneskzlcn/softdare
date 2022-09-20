package repository

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/entity"
	"time"
)

func (r *Repository) CreateComment(ctx context.Context, commentID, userID, postID, content string) (time.Time, error) {
	r.logger.Debugf("Create Comment Request Arrived To Repository With ID:%s", commentID)

	query := `
		INSERT INTO comments (id, user_id, post_id, content)
		VALUES ($1, $2, $3, $4) 
		RETURNING created_at;`

	row := r.db.QueryRowContext(ctx, query, commentID, userID, postID, content)
	var createdAt time.Time
	err := row.Scan(&createdAt)
	return createdAt, err

}

func (r *Repository) GetCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error) {
	r.logger.Debugf("Get Comments By Post ID Request Arrived To Repository With Post ID:%s", postID)
	query := `
		SELECT comments.id, comments.user_id, comments.post_id, comments.content, comments.like_count, comments.created_at, comments.updated_at, users.username 
		FROM comments
		INNER JOIN users ON comments.user_id = users.id 
		WHERE comments.post_id = $1   
		ORDER BY comments.id DESC;`

	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments := make([]*entity.Comment, 0)
	for rows.Next() {
		var c entity.Comment
		err = rows.Scan(
			&c.ID,
			&c.UserID,
			&c.PostID,
			&c.Content,
			&c.LikeCount,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.Username,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &c)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *Repository) AdjustCommentLikeCount(ctx context.Context, commentID string, adjustment int) (time.Time, error) {
	query := `
	UPDATE comments
	SET like_count = like_count+ $1, updated_at = now()
	RETURNING updated_at;`
	row := r.db.QueryRowContext(ctx, query, adjustment, commentID)
	var updatedAt time.Time
	if err := row.Scan(&updatedAt); err != nil {
		r.logger.Error(err)
		return time.Time{}, err
	}
	return updatedAt, nil
}
