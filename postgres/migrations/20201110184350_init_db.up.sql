SET TIMEZONE="utc";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id UUID DEFAULT uuid_generate_v4(),
  name VARCHAR(64) NOT NULL,
  email VARCHAR(320) UNIQUE,
  password_hash VARCHAR(60),
  deleted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE articles (
  id UUID DEFAULT uuid_generate_v4(),
  title VARCHAR(256) UNIQUE NOT NULL,
  content TEXT NOT NULL,
  thumbnail_url VARCHAR(256) NOT NULL,
  author UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT Now(),
  PRIMARY KEY (id),
  FOREIGN KEY(author) REFERENCES users(id)
);

CREATE TABLE images (
  id UUID DEFAULT uuid_generate_v4(),
  name VARCHAR(128),
  url VARCHAR(256) UNIQUE NOT NULL,
  img BYTEA NOT NULL,
  PRIMARY KEY (id)
);
