-- name: CreateUser :one
INSERT INTO users (id, username, locale, bio)
VALUES (@id, @username, @locale, @bio)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = @id;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = @username;

-- name: UpdateUser :one
UPDATE users
SET
    username   = COALESCE(sqlc.narg('username'), username),
    locale     = COALESCE(sqlc.narg('locale'), locale),
    bio        = COALESCE(sqlc.narg('bio'), bio),
    updated_at = now()
WHERE id = @id
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = @id;

-- name: UserExistsByID :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE id = @id
);

-- name: UsernameExists :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE username = @username
);
