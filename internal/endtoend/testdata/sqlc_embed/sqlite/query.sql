ATTACH 'baz.db' AS baz;

CREATE TABLE users (
    id integer PRIMARY KEY,
    name text NOT NULL,
    age integer
);

CREATE TABLE posts (
    id integer PRIMARY KEY,
    user_id integer NOT NULL
);

CREATE TABLE baz.users (
    id integer PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE user_links (
  owner_id integer NOT NULL,
  consumer_id integer NOT NULL,
  PRIMARY KEY (owner_id, consumer_id)
);

-- name: Only :one
SELECT sqlc.embed(users) FROM users;

-- name: WithAlias :one
SELECT sqlc.embed(u) FROM users AS u;

-- name: WithSubquery :many
SELECT sqlc.embed(users), (SELECT count(*) FROM users) AS total_count FROM users;

-- name: WithAsterisk :one
SELECT sqlc.embed(users), * FROM users;

-- name: Duplicate :one
SELECT sqlc.embed(users), sqlc.embed(users) FROM users;

-- name: Join :one
SELECT sqlc.embed(u), sqlc.embed(p) FROM posts AS p
INNER JOIN users AS u ON p.user_id = u.users.id;

-- name: WithSchema :one
SELECT sqlc.embed(bu) FROM baz.users AS bu;

-- name: WithCrossSchema :many
SELECT sqlc.embed(u), sqlc.embed(bu) FROM users AS u
INNER JOIN baz.users bu ON u.id = bu.id;

-- name: ListUserLink :many
SELECT
    sqlc.embed(owner),
    sqlc.embed(consumer)
FROM
    user_links
    INNER JOIN users AS owner ON owner.id = user_links.owner_id
    INNER JOIN users AS consumer ON consumer.id = user_links.consumer_id;