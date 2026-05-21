-- name: CreateUser :one
INSERT INTO Users (name, email) VALUES ($1, $2) RETURNING *;

-- name: GetUser :one
SELECT * FROM Users
WHERE uuid = $1;