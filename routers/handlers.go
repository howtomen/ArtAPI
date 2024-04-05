package router

import (
	"ArtAPI/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)


func getObjects(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Queryx("SELECT * FROM art_vault")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		objs := []models.ArtObject{}
		for rows.Next() {
			var obj models.ArtObject
			if err := rows.StructScan(&obj); err != nil {
				log.Fatal(err)
			}
			objs = append(objs, obj)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(objs)
	}
}

// get object by id
func getObject(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var obj models.ArtObject

		err := db.QueryRowx("SELECT * FROM art_vault WHERE id = $1", id).StructScan(&obj)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(obj)
	}
}

// create object
func createObject(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var obj models.ArtObject
		json.NewDecoder(r.Body).Decode(&obj)

		fmt.Println(obj)

		_, err := db.NamedExec("INSERT INTO art_vault (object_id,is_highlight,accession_year,primary_image,department,title,object_name,culture,period,artist_display_name,city,country) VALUES (:object_id,:is_highlight,:accession_year,:primary_image,:department,:title,:object_name,:culture,:period,:artist_display_name,:city,:country);", obj)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(obj)
	}
}

// update object
func updateObject(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var obj models.ArtObject
		json.NewDecoder(r.Body).Decode(&obj)

		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("UPDATE art_vault SET object_id = $1,is_highlight = $2,accession_year = $3,primary_image = $4,department = $5,title = $6,object_name = $7,culture = $8,period = $9,artist_display_name = $10,city = $11,country = $12 WHERE id = $13;", obj.ObjectID, obj.IsHighlight, obj.AccessionYear, obj.PrimaryImage, obj.Department, obj.Title, obj.ObjectName, obj.Culture, obj.Period, obj.ArtistDisplayName, obj.City, obj.Country, id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(obj)
	}
}

// delete object
func deleteObject(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var obj models.ArtObject
		err := db.QueryRowx("SELECT * FROM art_vault WHERE id = $1", id).StructScan(&obj)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			_, err := db.Exec("DELETE FROM art_vault WHERE id = $1", id)
			if err != nil {
				//todo: fix error handling
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}

		json.NewEncoder(w).Encode("Object deleted")
	}
}