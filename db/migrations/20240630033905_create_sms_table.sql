-- migrate:up
CREATE TABLE IF NOT EXISTS sms_messages (
    id TEXT PRIMARY KEY,
    inserted_at TEXT,
    body TEXT,
    sender TEXT,
    receiver TEXT
);

-- migrate:down
DROP TABLE sms;
