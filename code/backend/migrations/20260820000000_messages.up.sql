-- safe on populated tables: create canonical singleton row if missing
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS messages (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  content text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT ck_messages_content_non_empty CHECK (length(btrim(content)) > 0),
  CONSTRAINT ck_messages_content_single_line CHECK (content !~ '[\r\n]'),
  CONSTRAINT ck_messages_singleton CHECK (id = '00000000-0000-0000-0000-000000000001'::uuid)
);

INSERT INTO messages (id, content)
VALUES ('00000000-0000-0000-0000-000000000001', 'Hello Word')
ON CONFLICT (id) DO NOTHING;
