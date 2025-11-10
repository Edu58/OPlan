CREATE TABLE IF NOT EXISTS account_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    name VARCHAR(255) NOT NULL,
    active BOOLEAN DEFAULT FALSE,
    inserted_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE UNIQUE INDEX account_types_name_idx ON account_types (name);

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    msisdn VARCHAR(255),
    dob DATE,
    email_verified BOOLEAN DEFAULT false,
    msisdn_verified BOOLEAN DEFAULT false,
    active BOOLEAN DEFAULT false,

    account_type_id UUID REFERENCES account_types(id) ON DELETE RESTRICT,

    inserted_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE UNIQUE INDEX users_email_idx ON users (email);
CREATE UNIQUE INDEX users_username_idx ON users (username);
CREATE INDEX users_msisdn_idx ON users (msisdn);
