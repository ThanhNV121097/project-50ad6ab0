-- Create canonical singleton message row.
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

ALTER TABLE messages
    ADD COLUMN IF NOT EXISTS created_at timestamptz,
    ADD COLUMN IF NOT EXISTS updated_at timestamptz;

UPDATE messages
SET content = COALESCE(content, 'Hello Word'),
    created_at = COALESCE(created_at, now()),
    updated_at = COALESCE(updated_at, now())
WHERE content IS NULL OR created_at IS NULL OR updated_at IS NULL;

INSERT INTO messages (id, content)
VALUES ('00000000-0000-0000-0000-000000000001', 'Hello Word')
ON CONFLICT (id) DO NOTHING;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'ck_messages_content_non_empty'
    ) THEN
        ALTER TABLE messages
            ADD CONSTRAINT ck_messages_content_non_empty CHECK (length(btrim(content)) > 0);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'ck_messages_content_single_line'
    ) THEN
        ALTER TABLE messages
            ADD CONSTRAINT ck_messages_content_single_line CHECK (content !~ '[\r\n]');
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'ck_messages_singleton'
    ) THEN
        ALTER TABLE messages
            ADD CONSTRAINT ck_messages_singleton CHECK (id = '00000000-0000-0000-0000-000000000001'::uuid);
    END IF;
END $$;
