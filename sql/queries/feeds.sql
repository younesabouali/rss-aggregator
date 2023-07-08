
-- name: CreateFeed :one
INSERT INTO feeds (
  id, createdAt,updatedAt,name,url,user_id
)
VALUES ($1,$2,$3,$4, $5,$6)
returning *;

-- name: GetFeed :many
SELECT * FROM feeds LIMIT $1 OFFSET $2 ;

-- name: ListFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;

-- name: UpdateFeedFetchData :one
UPDATE  feeds 
SET last_fetched_at=$2
WHERE id=$1 
RETURNING *;
