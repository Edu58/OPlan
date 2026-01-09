CREATE TABLE "otp_store"
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    identifier VARCHAR(255) NOT NULL,
    value      VARCHAR(255) NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX identifier_idx ON otp_store (identifier);
CREATE INDEX expires_at_idx ON otp_store (expires_at);
