-- name: CreateUser :one
INSERT INTO users (
email,
first_name,
last_name,
password,
msisdn,
email_verified,
msisdn_verified,
active
) VALUES (
$1, 
$2, 
$3, 
$4, 
$5, 
$6, 
$7, 
$8
) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByIdForUpdate :one
SELECT * FROM users
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: GetUserByEmailForUpdate :one
SELECT * FROM users
WHERE email = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUserEmail :one
UPDATE users
SET email = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
