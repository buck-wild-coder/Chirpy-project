-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetToken :one
SELECT * FROM refresh_tokens where token = $1;

-- name: RevokeRefreshToken :exec
update refresh_tokens
set updated_at = $1, revoked_at = $2 
where token = $3;