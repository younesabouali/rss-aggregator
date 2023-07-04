
-- name: CreateFeed :one
INSERT INTO feeds (
  id, createdAt,updatedAt,name,url,user_id
)
VALUES ($1,$2,$3,$4, $5,$6)
returning *;

-- name: GetFeed :many
SELECT * FROM feeds LIMIT $1 OFFSET $2 ;
