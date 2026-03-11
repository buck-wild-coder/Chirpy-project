-- name: GetChirps :many
SELECT * from chirps ORDER by created_at;

-- name: GetChirpsAuthor :many
SELECT * from chirps where user_id = $1; 