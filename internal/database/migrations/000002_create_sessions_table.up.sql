CREATE TABLE "sessions" (
    user_id uuid PRIMARY KEY REFERENCES users (id),
    session_id VARCHAR(255) NOT NULL,
    client_ip varchar NOT NULL,
    is_blocked boolean NOT NULL DEFAULT false,
    expires_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE UNIQUE INDEX sessions_session_id_idx ON sessions (session_id);
CREATE INDEX sessions_client_ip_idx ON sessions (client_ip);
