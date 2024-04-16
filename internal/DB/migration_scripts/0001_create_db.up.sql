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


-- CREATE FUNCTION AND TRIGGER TO UPDATE SEARCH IDX TABLE BASED ON INSERT OR UPDATE 
-- TO art_vault
CREATE OR REPLACE FUNCTION update_search_idx() RETURNS TRIGGER AS $$
BEGIN
  IF TG_OP = 'INSERT' THEN
    INSERT INTO search_idx (id, search)
    VALUES (NEW.id, to_tsvector('simple', NEW.accession_year) || ' ' ||
                     to_tsvector('simple', NEW.department) || ' ' ||
                     to_tsvector('simple', NEW.title) || ' ' ||
                     to_tsvector('simple', NEW.object_name) || ' ' ||
                     to_tsvector('simple', NEW.culture) || ' ' ||
                     to_tsvector('simple', NEW.period) || ' ' ||
                     to_tsvector('simple', NEW.artist_display_name) || ' ' ||
                     to_tsvector('simple', NEW.city) || ' ' ||
                     to_tsvector('simple', NEW.country));
  ELSIF TG_OP = 'UPDATE' THEN
    UPDATE search_idx
    SET search = to_tsvector('simple', NEW.accession_year) || ' ' ||
                 to_tsvector('simple', NEW.department) || ' ' ||
                 to_tsvector('simple', NEW.title) || ' ' ||
                 to_tsvector('simple', NEW.object_name) || ' ' ||
                 to_tsvector('simple', NEW.culture) || ' ' ||
                 to_tsvector('simple', NEW.period) || ' ' ||
                 to_tsvector('simple', NEW.artist_display_name) || ' ' ||
                 to_tsvector('simple', NEW.city) || ' ' ||
                 to_tsvector('simple', NEW.country)
    WHERE id = NEW.id;
  ELSIF TG_OP = 'DELETE' THEN
    DELETE FROM search_idx
    WHERE uuid = OLD.id;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER art_row_updated
AFTER INSERT OR UPDATE OR DELETE ON art_vault
FOR EACH ROW
EXECUTE FUNCTION update_search_idx();