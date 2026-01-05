CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS "event_types" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    name CITEXT NOT NULL,
    description TEXT,
    active BOOLEAN DEFAULT true,

    inserted_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE UNIQUE INDEX event_types_name_idx ON event_types (name);

CREATE TABLE IF NOT EXISTS "events" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    name CITEXT NOT NULL,
    description TEXT,
    from_time TIMESTAMPTZ NOT NULL,
    to_time TIMESTAMPTZ NOT NULL,
    capacity INTEGER NOT NULL,
    policies_and_rules TEXT,
    min_age INTEGER NOT NULL,
    max_age INTEGER,

    age_restriction BOOLEAN DEFAULT false,
    public BOOLEAN DEFAULT true,
    require_ticket BOOLEAN DEFAULT false,

    event_type_id UUID REFERENCES event_types (id),

    inserted_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX events_min_age_idx ON events (min_age);
CREATE INDEX events_from_time_idx ON events (from_time);
CREATE INDEX events_age_restriction_idx ON events (age_restriction);
CREATE INDEX events_require_ticket_idx ON events (require_ticket);
CREATE INDEX events_event_type_id_idx ON events (event_type_id);
