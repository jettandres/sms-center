CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE sms_messages (
    id TEXT PRIMARY KEY,
    inserted_at TEXT,
    body TEXT,
    sender TEXT,
    receiver TEXT
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20240630033905');
