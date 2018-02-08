CREATE TABLE Card (
  id      SERIAL PRIMARY KEY,
  phrase_from varchar (255) NOT NULL,
  phrase_to   varchar (255) NOT NULL,
  lng_from    varchar (255) NOT NULL,
  lng_to      varchar (255) NOT NULL
);