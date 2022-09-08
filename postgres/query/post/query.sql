-- name: CreatePost :one

INSERT INTO posts (id, user_id, content)
VALUES (@id, @user_id, @content)
RETURNING  created_at;

