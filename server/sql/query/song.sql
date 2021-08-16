-- name: ListSongs :many
SELECT id, name, labels
FROM songs
ORDER BY id;

-- name: GetSong :one
SELECT id, name
FROM songs
WHERE id = $1;
