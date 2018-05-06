CREATE SEQUENCE IF NOT EXISTS quotes_sequence start 1;

CREATE TABLE IF NOT EXISTS quotes(
  id serial PRIMARY KEY,
  text varchar(2000) UNIQUE NOT NULL,
  owner varchar(250) NOT NULL,
  date date NOT NULL,
  created_at timestamp NOT NULL
);
