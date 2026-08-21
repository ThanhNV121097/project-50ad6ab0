-- Create canonical singleton message row.
CREATE EXTENSION IF NOT EXISTS pgcrypto;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.tables
        WHERE table_schema = current_schema()
          AND table_name = 'messages'
    ) THEN
        CREATE TABLE messages (
            id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
            content text NOT NULL,
            created_at timestamptz NOT NULL DEFAULT now(),
            updated_at timestamptz NOT NULL DEFAULT now(),
            CONSTRAINT ck_messages_content_non_empty CHECK (length(btrim(content)) > 0),
            CONSTRAINT ck_messages_content_single_line CHECK (content !~ '[\r\n]'),
            CONSTRAINT ck_messages_singleton CHECK (id = '00000000-0000-0000-0000-000000000001'::uuid)
        );
    ELSE
        DELETE FROM messages;

        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint
            WHERE conname = 'ck_messages_content_non_empty'
              AND conrelid = 'messages'::regclass
        ) THEN
            ALTER TABLE messages
                ADD CONSTRAINT ck_messages_content_non_empty CHECK (length(btrim(content)) > 0);
        END IF;

        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint
            WHERE conname = 'ck_messages_content_single_line'
              AND conrelid = 'messages'::regclass
        ) THEN
            ALTER TABLE messages
                ADD CONSTRAINT ck_messages_content_single_line CHECK (content !~ '[\r\n]');
        END IF;

        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint
            WHERE conname = 'ck_messages_singleton'
              AND conrelid = 'messages'::regclass
        ) THEN
            ALTER TABLE messages
                ADD CONSTRAINT ck_messages_singleton CHECK (id = '00000000-0000-0000-0000-000000000001'::uuid);
        END IF;
    END IF;
END $$;

INSERT INTO messages (id, content)
VALUES ('00000000-0000-0000-0000-000000000001', 'Hello Word')
ON CONFLICT (id) DO NOTHING;

