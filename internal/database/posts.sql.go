// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id,created_at, updated_at,title, url, description, published_at,feed_id) VALUES ( $1,$2,$3,$4,$5,$6,$7,$8 )
RETURNING id, created_at, updated_at, title, url, description, published_at, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt sql.NullTime
	FeedID      uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const getFollowedPosts = `-- name: GetFollowedPosts :many
SELECT feed_follows.id, feed_follows.created_at, feed_follows.updated_at, user_id, feed_follows.feed_id, posts.id, posts.created_at, posts.updated_at, title, url, description, published_at, posts.feed_id from feed_follows INNER JOIN posts ON feed_follows.feed_id=posts.feed_id WHERE feed_follows.user_id=$1 LIMIT $2 OFFSET $3
`

type GetFollowedPostsParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

type GetFollowedPostsRow struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uuid.UUID
	FeedID      uuid.UUID
	ID_2        uuid.UUID
	CreatedAt_2 time.Time
	UpdatedAt_2 time.Time
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt sql.NullTime
	FeedID_2    uuid.UUID
}

func (q *Queries) GetFollowedPosts(ctx context.Context, arg GetFollowedPostsParams) ([]GetFollowedPostsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFollowedPosts, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFollowedPostsRow
	for rows.Next() {
		var i GetFollowedPostsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
