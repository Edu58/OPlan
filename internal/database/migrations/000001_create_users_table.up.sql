CREATE TABLE IF NOT EXISTS "users" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    email VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    msisdn VARCHAR(255),
    email_verified BOOLEAN DEFAULT false,
    msisdn_verified BOOLEAN DEFAULT false,
    active BOOLEAN DEFAULT false,

    inserted_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE UNIQUE INDEX users_email_idx ON users (email);
CREATE INDEX users_msisdn_idx ON users (msisdn);
CREATE INDEX users_active_idx ON users (active);
