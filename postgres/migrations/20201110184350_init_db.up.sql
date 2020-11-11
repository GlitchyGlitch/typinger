CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE role AS ENUM ('ADMIN', 'AUTHOR');


CREATE TABLE users (
  id UUID DEFAULT uuid_generate_v4(),
  name VARCHAR(55) NOT NULL,
  email VARCHAR(255) UNIQUE,
  role ROLE,
  password VARCHAR(60),
  active BOOLEAN DEFAULT true NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE articles (
  id UUID DEFAULT uuid_generate_v4(),
  title VARCHAR(255) NOT NULL,
  content TEXT UNIQUE NOT NULL,
  thumbnail_url VARCHAR(255) NOT NULL,
  author UUID NOT NULL,
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