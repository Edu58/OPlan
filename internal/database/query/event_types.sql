-- name: ListEventTypes :many
SELECT * FROM event_types
WHERE active;

-- name: CreateEventType :one
INSERT INTO event_types (
    name,
    description,
    active,
    inserted_at,
    updated_at
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetEventTypeById :one
SELECT * FROM event_types
WHERE id = $1
AND active
LIMIT 1;

-- name: GetEventTypeByName :one
SELECT * FROM event_types
WHERE name = $1
AND active
LIMIT 1;

-- name: UpdateEventTypeById :one
UPDATE event_types
SET
  name = $2,
  description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteEventType :exec
DELETE FROM event_types
WHERE id = $1;
