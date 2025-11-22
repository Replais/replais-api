-- users
CREATE TABLE IF NOT EXISTS users (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email       CITEXT NOT NULL UNIQUE,
  name        TEXT NOT NULL,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- personas (global + user-level overrides)
CREATE TABLE IF NOT EXISTS personas (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id         UUID REFERENCES users(id) ON DELETE CASCADE,

  slug            CITEXT NOT NULL,
  label           TEXT NOT NULL,
  description     TEXT,
  prompt_template TEXT,
  is_default      BOOLEAN NOT NULL DEFAULT FALSE,

  created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),

  UNIQUE (user_id, slug)
);

-- platform-level user settings
CREATE TABLE IF NOT EXISTS user_settings (
  id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id            UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  platform           TEXT NOT NULL,  -- 'whatsapp', 'gmail', etc.
  default_persona_id UUID REFERENCES personas(id),
  global_instructions TEXT NOT NULL DEFAULT '',
  extra              JSONB NOT NULL DEFAULT '{}'::jsonb,
  created_at         TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at         TIMESTAMPTZ NOT NULL DEFAULT now(),

  UNIQUE (user_id, platform)
);