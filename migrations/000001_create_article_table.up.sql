CREATE EXTENSION "uuid-ossp";

CREATE TABLE article
(
  id UUID PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
  title TEXT NOT NULL,
  author TEXT,
  published_at TIMESTAMP WITHOUT TIME ZONE,
  publisher TEXT,
  cover_url TEXT,
  overview TEXT
);
