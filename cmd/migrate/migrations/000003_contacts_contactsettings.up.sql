CREATE TABLE IF NOT EXISTS contacts (
  id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id             UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  platform            TEXT NOT NULL,
  platform_contact_key TEXT NOT NULL,   -- e.g. "whatsapp|+91xxxx" or thread id
  display_name        TEXT NOT NULL,
  is_group            BOOLEAN NOT NULL DEFAULT FALSE,
  created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at          TIMESTAMPTZ NOT NULL DEFAULT now(),

  UNIQUE (user_id, platform, platform_contact_key)
);

CREATE TABLE IF NOT EXISTS contact_settings (
  id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id      UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  contact_id   UUID NOT NULL REFERENCES contacts(id) ON DELETE CASCADE,
  persona_id   UUID REFERENCES personas(id),
  instructions TEXT NOT NULL DEFAULT '',
  extra        JSONB NOT NULL DEFAULT '{}'::jsonb,
  created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),

  -- one settings row per contact per user
  UNIQUE (user_id, contact_id)
);