-- name: ListAccountTypes :many
SELECT * FROM account_types
ORDER BY inserted_at DESC;

-- name: GetAccountTypeById :one
SELECT * FROM account_types
WHERE id = $1 LIMIT 1;

-- name: GetAccountTypeByName :one
SELECT * FROM account_types
WHERE name = $1 LIMIT 1;

-- name: CreateAccountType :one
INSERT INTO account_types
(name, active)
VALUES($1, $2)
RETURNING *;

-- name: UpdateAccountTypeByID :one
UPDATE account_types
SET name=$2, active=$2
WHERE id = $1
RETURNING *;

-- name: DeleteAccountType :one
DELETE FROM account_types
WHERE id = $1 OR name = $1
RETURNING *;