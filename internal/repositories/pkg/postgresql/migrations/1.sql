-- +migrate Up
CREATE TABLE IF NOT EXISTS events (
  "id" serial primary key not null,
  "realm_id" text NOT NULL,
  "event_id" text NOT NULL,
  "type" text NOT NULL,
  "body" bytea NOT NULL,
  "inserted_at" timestamp(6) NOT NULL DEFAULT statement_timestamp()
);

-- +migrate Down
DROP TABLE events;