-- name: CreatePost :one

INSERT INTO posts (id, user_id, content)
VALUES (@id, @user_id, @content)
RETURNING  created_at;

-- name: Posts :many
SELECT posts.*, users.username
FROM posts
INNER JOIN users ON posts.user_id = users.id
ORDER BY posts.id;

-- name: PostById :one
SELECT posts.*, users.username
FROM posts
INNER JOIN users ON posts.user_id = users.id
WHERE posts.id = @post_id;