-- name: CreateComment :one

INSERT INTO comments (id, user_id, post_id, content)
VALUES (@id, @user_id, @post_id, @content)
RETURNING created_at;


-- name: GetCommentsByPostID :many

SELECT comments.*,users.username
FROM comments
INNER JOIN users ON comments.user_id = users.id
WHERE comments.post_id = @post_id
ORDER BY comments.id DESC;
