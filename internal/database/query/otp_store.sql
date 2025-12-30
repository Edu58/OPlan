-- name: CreateOTP :one
INSERT INTO otp_store (id,
                       identifier,
                       value,
                       expires_at)
VALUES ($1, $2, $3, $4)
ON CONFLICT (identifier)
    DO UPDATE SET value      = EXCLUDED.value,
                  expires_at = EXCLUDED.expires_at
RETURNING *;

-- name: GetOTP :one
SELECT *
FROM otp_store
WHERE identifier = $1
LIMIT 1;

-- name: UpdateOTP :one
UPDATE otp_store
SET value     = $2,
    expires_at=$3
WHERE identifier = $1
RETURNING *;

-- name: DeleteOTP :exec
DELETE
FROM otp_store
WHERE identifier = $1;