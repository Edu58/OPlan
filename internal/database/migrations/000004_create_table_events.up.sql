CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS "event_types" (
    id  uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    name CITEXT NOT NULL,
    description TEXT,
    active BOOLEAN DEFAULT true,

    inserted_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE UNIQUE INDEX event_types_name_idx ON event_types (name);

CREATE TABLE IF NOT EXISTS "events" (
    id  uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    public BOOLEAN DEFAULT true,
    require_ticket BOOLEAN DEFAULT false,

    event_type_id UUID REFERENCES event_types (id)
);

CREATE INDEX events_public_idx ON events (public);
CREATE INDEX events_require_ticket_idx ON events (require_ticket);
CREATE INDEX events_event_type_id_idx ON events (event_type_id);
