

CREATE TABLE IF NOT EXISTS users (
  id serial,
  name text,
  email text not null,
  password text not null
);

CREATE TABLE IF NOT EXISTS wechat (
  id serial,
  name text not null,
  wechat_id text not null
);