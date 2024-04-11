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

-- CREATE TABLE THAT WILL ALLOW US TO SEARCH RECORDS 
CREATE TABLE IF NOT EXISTS search_idx
(
  id uuid PRIMARY KEY NOT NULL,
  search tsvector NOT NULL
);

CREATE FUNCTION update_search_idx() RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO search_idx (id, search)
    SELECT NEW.id, to_tsvector('simple', NEW.accession_year) || ' ' ||
                   to_tsvector('simple', NEW.department) || ' ' ||
                   to_tsvector('simple', NEW.title) || ' ' ||
                   to_tsvector('simple', NEW.object_name) || ' ' ||
                   to_tsvector('simple', NEW.culture) || ' ' ||
                   to_tsvector('simple', NEW.period) || ' ' ||
                   to_tsvector('simple', NEW.artist_display_name) || ' ' ||
                   to_tsvector('simple', NEW.city) || ' ' ||
                   to_tsvector('simple', NEW.country);
  RETURN NEW;
END; 
$$ LANGUAGE plpgsql;

CREATE TRIGGER art_row_added
AFTER INSERT ON art_vault
FOR EACH ROW
EXECUTE FUNCTION update_search_idx();