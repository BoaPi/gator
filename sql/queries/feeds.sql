-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ( 
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follow (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE $1 = url; 

-- name: GetFeedFollowsForUser :many
SELECT
    feed_follow.id,
    feed_follow.created_at,
    feed_follow.updated_at,
    feed_follow.user_id,
    feed_follow.feed_id,
    users.name AS user_name,
    feeds.name AS feed_name
FROM feed_follow
INNER JOIN users ON feed_follow.user_id = users.id
INNER JOIN feeds ON feed_follow.feed_id = feeds.id
WHERE feed_follow.user_id = $1;
