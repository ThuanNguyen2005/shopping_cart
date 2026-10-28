-- name: CreateUser :one
INSERT INTO users (user_name, user_email, user_password, user_age, user_status, user_level) 
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: UpdateUser :one
UPDATE users 
SET 
    user_name = COALESCE(sqlc.narg(user_name),user_name),
    user_password = COALESCE(sqlc.narg(user_password),user_password),
    user_age = COALESCE(sqlc.narg(user_age),user_age),
    user_status = COALESCE(sqlc.narg(user_status),user_status),
    user_level = COALESCE(sqlc.narg(user_level),user_level)
WHERE 
    user_uuid=sqlc.arg(user_uuid) AND delete_at IS NULL
RETURNING *;

-- name: SoftDeleteUser :one
UPDATE users
SET 
    delete_at=NOW()
WHERE 
    user_uuid=sqlc.arg(user_uuid) AND delete_at IS NULL
RETURNING *;

-- name: RecoveUser :one
UPDATE users
SET 
    delete_at=NULL
WHERE 
    user_uuid=sqlc.arg(user_uuid) AND delete_at IS NOT NULL
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE 
    user_uuid=sqlc.arg(user_uuid)
RETURNING *;

-- name: ListUsersIDASC :many
SELECT *
FROM users
WHERE delete_at IS NULL
AND(
    sqlc.narg(search)::TEXT IS NULL
    OR sqlc.narg(search)::TEXT=''
    OR user_name ILIKE '%'|| sqlc.narg(search) ||'%'
    OR user_email ILIKE '%'|| sqlc.narg(search) ||'%'
)
ORDER By user_id
LIMIT $1 OFFSET $2;

-- name: ListUsersIDDESC :many
SELECT *
FROM users
WHERE delete_at IS NULL
AND(
    sqlc.narg(search)::TEXT IS NULL
    OR sqlc.narg(search)::TEXT=''
    OR user_name ILIKE '%'|| sqlc.narg(search) ||'%'
    OR user_email ILIKE '%'|| sqlc.narg(search) ||'%'
)
ORDER By user_id DESC
LIMIT $1 OFFSET $2;

-- name: ListUsersCreateAtASC :many
SELECT *
FROM users
WHERE delete_at IS NULL
AND(
    sqlc.narg(search)::TEXT IS NULL
    OR sqlc.narg(search)::TEXT=''
    OR user_name ILIKE '%'|| sqlc.narg(search) ||'%'
    OR user_email ILIKE '%'|| sqlc.narg(search) ||'%'
)
ORDER By create_at
LIMIT $1 OFFSET $2;

-- name: ListUsersCreateAtDESC :many
SELECT *
FROM users
WHERE delete_at IS NULL
AND(
    sqlc.narg(search)::TEXT IS NULL
    OR sqlc.narg(search)::TEXT=''
    OR user_name ILIKE '%'|| sqlc.narg(search) ||'%'
    OR user_email ILIKE '%'|| sqlc.narg(search) ||'%'
)
ORDER By create_at DESC
LIMIT $1 OFFSET $2;

-- name: UserTotal :one
SELECT COUNT(1)
FROM users
WHERE delete_at IS NULL
AND(
    sqlc.narg(search)::TEXT IS NULL
    OR sqlc.narg(search)::TEXT=''
    OR user_name ILIKE '%'|| sqlc.narg(search) ||'%'
    OR user_email ILIKE '%'|| sqlc.narg(search) ||'%'
);

