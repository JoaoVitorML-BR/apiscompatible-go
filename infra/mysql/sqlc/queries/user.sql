-- name: CreateUser :exec
INSERT INTO users(name, password) VALUES(?, ?);

-- name: FindUsers :many
SELECT id, name, password FROM users;

-- name: FindUserByID :one
SELECT id, name, password FROM users WHERE id = ?;

-- name: FindUserByName :many
SELECT id, name FROM users WHERE name LIKE ?;

-- name: FindUserByNameToLogin :one
SELECT id, name, password FROM users WHERE name = ?;

-- name: UpdateUser :exec
UPDATE users SET name = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;