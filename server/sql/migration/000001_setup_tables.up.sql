CREATE TABLE IF NOT EXISTS songs (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL,
  labels JSONB DEFAULT '[]'::jsonb,
  data JSONB DEFAULT '{}'::jsonb
);

