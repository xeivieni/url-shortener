CREATE TABLE urls (
  id          SERIAL PRIMARY KEY,
  shorturl    VARCHAR(10) NOT NULL,
  longurl     VARCHAR(140),
  hits	      INTEGER NOT NULL DEFAULT 0
);
