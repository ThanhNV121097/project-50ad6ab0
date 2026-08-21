-- Create canonical singleton message row.
CREATE EXTENSION IF NOT EXISTS pgcrypto;

ALTER TABLE IF EXISTS messages
    ADD COLUMN IF NOT EXISTS id uuid,
    ADD COLUMN IF NOT EXISTS content text,
    ADD COLUMN IF NOT EXISTS created_at timestamptz,
    ADD COLUMN IF NOT EXISTS updated_at timestamptz;

CREATE TABLE IF NOT EXISTS messages (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    content text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT ck_messages_content_non_empty CHECK (length(btrim(content)) > 0),
    CONSTRAINT ck_messages_content_single_line CHECK (content !~ '[\r\n]'),
    CONSTRAINT ck_messages_singleton CHECK (id = '00000000-0000-0000-0000-000000000001'::uuid)
);

ALTER TABLE IF EXISTS messages
    ALTER COLUMN id SET DEFAULT gen_random_uuid(),
    ALTER COLUMN content SET DEFAULT '',
    ALTER COLUMN created_at SET DEFAULT now(),
    ALTER COLUMN updated_at SET DEFAULT now();

UPDATE messages
SET id = COALESCE(id, '00000000-0000-0000-0000-000000000001'::uuid),
    content = COALESCE(content, 'Hello Word'),
    created_at = COALESCE(created_at, now()),
    updated_at = COALESCE(updated_at, now())
WHERE id IS NULL OR content IS NULL OR created_at IS NULL OR updated_at IS NULL;

UPDATE messages
SET content = 'Hello Word'
WHERE id = '00000000-0000-0000-0000-000000000001'::uuid AND btrim(content) = '';

ALTER TABLE messages
    ALTER COLUMN id SET NOT NULL,
    ALTER COLUMN content SET NOT NULL,
    ALTER COLUMN created_at SET NOT NULL,
    ALTER COLUMN updated_at SET NOT NULL;

ALTER TABLE messages
    ADD CONSTRAINT IF NOT EXISTS pk_messages PRIMARY KEY (id),
    ADD CONSTRAINT IF NOT EXISTS ck_messages_content_non_empty CHECK (length(btrim(content)) > 0),
    ADD CONSTRAINT IF NOT EXISTS ck_messages_content_single_line CHECK (content !~ '[\r\n]'),
    ADD CONSTRAINT IF NOT EXISTS ck_messages_singleton CHECK (id = '00000000-0000-0000-0000-000000000001'::uuid);

INSERT INTO messages (id, content)
VALUES ('00000000-0000-0000-0000-000000000001', 'Hello Word')
ON CONFLICT (id) DO NOTHING;
