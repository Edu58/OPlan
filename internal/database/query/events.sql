-- name: ListEvents :many
SELECT events.*, sqlc.embed(event_types) FROM events
JOIN event_types ON event_types.id = events.event_type_id
LIMIT $1 OFFSET $2;

-- name: CreateEvent :one
INSERT INTO events (
    name,
    from_time,
    to_time,
    capacity,
    policies_and_rules,
    min_age,
    max_age,
    age_restriction,
    public,
    require_ticket,
    event_type_id,
    inserted_at,
    updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING *;

-- name: GetEventById :one
SELECT * FROM events
WHERE id = $1
LIMIT 1;

-- name: GetEventByName :one
SELECT * FROM events
WHERE name = $1
LIMIT 1;

-- name: UpdateEventById :one
UPDATE events
SET
name = $2,
from_time = $3,
to_time = $4,
capacity = $5,
policies_and_rules = $6,
min_age = $7,
max_age = $8,
age_restriction = $9,
public = $10,
require_ticket = $11,
event_type_id = $11
WHERE id = $1
RETURNING *;

-- name: DeleteEventById :exec
DELETE FROM events
WHERE id = $1;
