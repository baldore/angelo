-- name: ListSongs :many
SELECT id, name, labels
FROM songs
ORDER BY id;

-- name: GetSong :one
SELECT id, name
FROM songs
WHERE id = $1;

-- name: CreateSong :one
INSERT INTO songs (name)
VALUES ($1)
RETURNING *;

-- name: UpdateSong :exec
UPDATE songs
SET labels = $1
WHERE id = $2;

-- name: DeleteSong :exec
DELETE FROM songs
WHERE id = $1;
