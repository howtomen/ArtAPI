CREATE TABLE art_vault
(
  "id" SERIAL PRIMARY KEY,
  "object_id" bigint NOT NULL,
  "is_highlight" boolean DEFAULT false NOT NULL,
  "accession_year" text,
  "primary_image" text,
  "department" text,
  "title" text,
  "object_name" text,
  "culture" text,
  "period" text,
  "artist_display_name" text,
  "city" text,
  "country" text
);


