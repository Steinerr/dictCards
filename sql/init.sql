CREATE TABLE page (
  id      SERIAL PRIMARY KEY,
  title   varchar(255) UNIQUE NOT NULL,
  body    text
);

INSERT INTO page (title, body) VALUES
  ('test', 'TEST PAGE'),
  ('111', 'NUMBERS')
;