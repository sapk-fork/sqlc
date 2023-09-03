CREATE TABLE users (
  id INT UNSIGNED NOT NULL,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE user_links (
  owner_id INT UNSIGNED NOT NULL,
  consumer_id INT UNSIGNED NOT NULL,
  PRIMARY KEY (owner_id, consumer_id)
);

-- name: ListUserRelation :many
SELECT
    sqlc.embed(owner),
    sqlc.embed(consumer)
FROM
    user_links
    INNER JOIN users AS owner ON owner.id = user_links.owner_id
    INNER JOIN users AS consumer ON consumer.id = user_links.consumer_id;