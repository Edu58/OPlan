CREATE TABLE "sessions" (
    user_id UUID PRIMARY KEY REFERENCES users (id),
    session_id UUID DEFAULT gen_random_uuid(),
    client_ip varchar(255),
    is_blocked boolean NOT NULL DEFAULT false,
    expires_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE UNIQUE INDEX sessions_user_id_session_id_idx ON sessions (user_id, session_id);
CREATE INDEX sessions_client_ip_idx ON sessions (client_ip);
