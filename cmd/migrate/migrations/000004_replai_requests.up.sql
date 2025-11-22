CREATE TABLE IF NOT EXISTS replai_requests (
  id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id          UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  contact_id       UUID REFERENCES contacts(id) ON DELETE SET NULL,
  platform         TEXT NOT NULL,
  context          JSONB NOT NULL,
  user_context     JSONB NOT NULL,
  ui_state         JSONB NOT NULL,
  model_name       TEXT NOT NULL,
  prompt_tokens    INT,
  completion_tokens INT,
  latency_ms       INT,
  reply_text       TEXT,
  created_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);