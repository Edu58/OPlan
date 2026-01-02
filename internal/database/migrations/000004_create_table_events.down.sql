DROP INDEX IF EXISTS events_event_type_id_idx ;
DROP INDEX IF EXISTS events_require_ticket_idx ;
DROP INDEX IF EXISTS events_public_idx;
DROP TABLE IF EXISTS events;

DROP INDEX IF EXISTS event_types_name_idx;
DROP TABLE IF EXISTS event_types;