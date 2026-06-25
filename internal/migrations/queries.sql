-- name: CreateUser :one
INSERT INTO users(username, email, password, created, updated)
VALUES ($1, $2, $3, $4, $5)
  RETURNING id, username, email, created, updated;

-- name: GetUser :one
SELECT id, username, email, created, updated
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, username, email, created, updated
FROM users
ORDER BY id;

-- name: CreateItem :one
INSERT INTO items(title, price, user_id, created, updated)
VALUES ($1, $2, $3, $4, $5)
  RETURNING id, title, price, user_id, created, updated;

-- name: ListItems :many
SELECT id, title, price, user_id, created, updated
FROM items
ORDER BY id;
