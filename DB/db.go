package database

import (
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func NewDbConnection() *sqlx.DB {
	db, err := sqlx.Open("postgres", "dbname=go_db user=postgres password=postgres port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	//check if table exists, if not create table. 
	tableName := "art_vault"
	exists, err := tableExists(db, tableName)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		err = createTable(db)
		if err != nil {
			log.Fatal(err)
		}
	}
	return db
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
