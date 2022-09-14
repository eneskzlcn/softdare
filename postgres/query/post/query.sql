-- name: CreatePost :one

INSERT INTO posts (id, user_id, content)
VALUES (@id, @user_id, @content)
RETURNING  created_at;

-- name: Posts :many
SELECT posts.*, users.username
FROM posts
INNER JOIN users ON posts.user_id = users.id
WHERE
    CASE
        WHEN @user_id <> '' THEN users.id = @user_id
        ELSE true
    END
ORDER BY posts.id DESC;

-- name: PostById :one
SELECT posts.*, users.username
FROM posts
INNER JOIN users ON posts.user_id = users.id
WHERE posts.id = @post_id;

-- name: UpdatePost :one

UPDATE posts
SET comment_count = comment_count + @comment_count_increase_amount, updated_at = now()
WHERE id = @post_id
RETURNING updated_at;
