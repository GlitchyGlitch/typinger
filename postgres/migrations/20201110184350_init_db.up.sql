SET TIMEZONE="utc";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id UUID DEFAULT uuid_generate_v4(),
  name VARCHAR(55) NOT NULL,
  email VARCHAR(255) UNIQUE,
  password_hash VARCHAR(60),
  deleted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE articles (
  id UUID DEFAULT uuid_generate_v4(),
  title VARCHAR(255) UNIQUE NOT NULL,
  content TEXT NOT NULL,
  thumbnail_url VARCHAR(255) NOT NULL,
  author UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT Now(),
  PRIMARY KEY (id),
  FOREIGN KEY(author) REFERENCES users(id)
);

CREATE TABLE images (
  id UUID DEFAULT uuid_generate_v4(),
  name VARCHAR(127),
  url VARCHAR(255) UNIQUE NOT NULL,
  img BYTEA NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE settings (
  id UUID DEFAULT uuid_generate_v4(),
  key VARCHAR(255) UNIQUE NOT NULL,
  value TEXT NOT NULL,
  PRIMARY KEY (id)
);