-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id, last_fetched_at)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7
)
RETURNING *;

-- name: ListAllFeeds :many
SELECT feeds.*, users.name AS user_name
FROM feeds
INNER JOIN users 
ON feeds.user_id = users.id;

-- name: GetFeed :one
SELECT * FROM feeds 
WHERE url = $1 LIMIT 1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = $1,
updated_at = $1
WHERE id = $2
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;
