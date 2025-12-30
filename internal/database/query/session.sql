-- name: CreateSession :one
INSERT INTO sessions (
    user_id,
    session_id,
    client_ip,
    is_blocked,
    expires_at
) VALUES (
  $1, $2, $3, $4, $5
) ON CONFLICT (user_id, session_id)
    DO UPDATE SET session_id = EXCLUDED.session_id,
                  expires_at = EXCLUDED.expires_at
  RETURNING *;

-- name: GetSessionByUserId :one
SELECT * FROM sessions
WHERE user_id = $1 LIMIT 1;

-- name: GetSessionBySessionId :one
SELECT * FROM sessions
WHERE session_id = $1 LIMIT 1;

-- name: UpdateSessionIsBlocked :one
UPDATE sessions
SET
  is_blocked = $2
WHERE session_id = $1
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE session_id = $1;