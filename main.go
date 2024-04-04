package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/jmoiron/sqlx"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type ArtObject struct {
	ID int `json:"id" db:"id"` 
	ObjectID int `json:"object_id" db:"object_id"`
	IsHighlight bool `json:"is_highlight" db:"is_highlight"`
	AccessionYear string `json:"accession_year" db:"accession_year"`
	PrimaryImage string `json:"primary_image" db:"primary_image"`
	Department string `json:"department" db:"department"`
	Title string `json:"title" db:"title"`
	ObjectName string `json:"object_name" db:"object_name"`
	Culture string `json:"culture" db:"culture"`
	Period string `json:"period" db:"period"`
	ArtistDisplayName string `json:"artist_display_name" db:"artist_display_name"`
	City string `json:"city" db:"city"`
	Country string `json:"country" db:"country"`
}

func main() {
	//open connection to DB
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	//check if table exists, if not create table. 
	tableName := "art_vault"
	exists, err := tableExists(db, tableName)
	if err != nil {
		log.Println(err)
	}

	if !exists {
		err = createTable(db)
		if err != nil {
			log.Fatal(err)
		}
	}

	//create router
	router := mux.NewRouter()
	router.HandleFunc("/objects",getObjects(db)).Methods("GET")
	router.HandleFunc("/objects/{id}", getObject(db)).Methods("GET")
	router.HandleFunc("/objects", createObject(db)).Methods("POST")
	router.HandleFunc("/objects/{id}", updateObject(db)).Methods("PUT")
	router.HandleFunc("/objects/{id}", deleteObject(db)).Methods("DELETE")
	//start server
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}

func tableExists(db *sqlx.DB, table string) (bool, error) {
	query := fmt.Sprintf(`
        SELECT EXISTS (
            SELECT FROM information_schema.tables 
            WHERE  table_schema = 'public'
            AND    table_name   = '%s'
        );
    `, table)
	var exists bool 
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func createTable(db *sqlx.DB) error{
	sqlLines, err := readSQLFile("create_db.sql")
	if err != nil {
		log.Fatal(err)
	}

	_,err = db.Exec(sqlLines)
	return err
}

func readSQLFile(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w,r)
	})
}

//get all objects
func getObjects(db *sqlx.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		rows, err := db.Queryx("SELECT * FROM art_vault")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		objs := []ArtObject{}
		for rows.Next() {
			var obj ArtObject
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
//get object by id
func getObject(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var obj ArtObject
		
		err := db.QueryRowx("SELECT * FROM art_vault WHERE id = $1", id).StructScan(&obj)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(obj)
	}
}
//create object
func createObject(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var obj ArtObject
		json.NewDecoder(r.Body).Decode(&obj)

		fmt.Println(obj)

		_,err := db.NamedExec("INSERT INTO art_vault (object_id,is_highlight,accession_year,primary_image,department,title,object_name,culture,period,artist_display_name,city,country) VALUES (:object_id,:is_highlight,:accession_year,:primary_image,:department,:title,:object_name,:culture,:period,:artist_display_name,:city,:country);", obj)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(obj)
	}
}
//update object
func updateObject(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var obj ArtObject
		json.NewDecoder(r.Body).Decode(&obj)

		vars := mux.Vars(r)
		id := vars["id"]

		_,err := db.Exec("UPDATE art_vault SET object_id = $1,is_highlight = $2,accession_year = $3,primary_image = $4,department = $5,title = $6,object_name = $7,culture = $8,period = $9,artist_display_name = $10,city = $11,country = $12 WHERE id = $13;",obj.ObjectID, obj.IsHighlight,obj.AccessionYear,obj.PrimaryImage,obj.Department,obj.Title,obj.ObjectName,obj.Culture,obj.Period,obj.ArtistDisplayName,obj.City,obj.Country,id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(obj)
	}
}
//delete object
func deleteObject(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var obj ArtObject
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