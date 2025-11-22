-- Enable pgcrypto for gen_random_uuid()
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Enable citext for case-insensitive text
CREATE EXTENSION IF NOT EXISTS citext;