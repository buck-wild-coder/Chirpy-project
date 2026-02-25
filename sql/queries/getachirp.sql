-- name: GetAChirp :one
SELECT * from chirps where id = $1;