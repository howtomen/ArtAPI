-- naming 0001 since the migration library uses that 
-- naming convention to dictate execution order
-- this is also why there is a .up and .down
CREATE TABLE IF NOT EXISTS art_vault
(
  "id" uuid PRIMARY KEY,
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

