DROP INDEX CONCURRENTLY IF EXISTS  sessions_session_id_idx;
DROP INDEX CONCURRENTLY IF EXISTS sessions_client_ip_idx;

DROP TABLE IF EXISTS "sessions";
