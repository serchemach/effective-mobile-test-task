CREATE TABLE IF NOT EXISTS songs(
   song_id serial PRIMARY KEY,
   name text NOT NULL,
   artist text NOT NULL,
   release_date text NOT NULL,
   lyrics text NOT NULL,
   link text NOT NULL
);
