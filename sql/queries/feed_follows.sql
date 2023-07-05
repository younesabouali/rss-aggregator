-- name: FollowFeed :one
INSERT INTO feed_follows  (id,created_at,updated_at,user_id,feed_id) VALUES ($1,$2,$3,$4,$5)
RETURNING * ;

-- name: DeleteFeedFollow :one
DELETE FROM feed_follows WHERE feed_id =$1 AND user_id=$2
RETURNING *;

-- name: GetFollowedFeed :many
SELECT * FROM feed_follows RIGHT JOIN feeds ON feed_follows.feed_id = feeds.id WHERE feed_follows.user_id=$1 LIMIT $2 OFFSET $3;
