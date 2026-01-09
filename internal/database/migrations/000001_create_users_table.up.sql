CREATE TABLE IF NOT EXISTS "users" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    email VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    msisdn VARCHAR(255),
    email_verified BOOLEAN NOT NULL DEFAULT False,
    msisdn_verified BOOLEAN NOT NULL DEFAULT False,
    active BOOLEAN NOT NULL DEFAULT False,

    inserted_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX users_email_idx ON users (email);
CREATE INDEX users_msisdn_idx ON users (msisdn);
CREATE INDEX users_active_idx ON users (active);
