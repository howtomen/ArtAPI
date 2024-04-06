package db

import (
	"context"
	"fmt"
	// "log"
	"os"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	Client *sqlx.DB
}

func NewDbConnection() (*Database, error) {
	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)
	db, err := sqlx.Connect("postgres", connectionStr)
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to the database: %w", err)
	}

	return &Database{
		Client: db,
	}, nil 

	// //check if table exists, if not create table.
	// exists, err := tableExists(db, "art_vault")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if !exists {
	// 	err = createTable(db)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// return db
}

//Adding ping function so that later we can health check the conn to DB easily
func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}

// func tableExists(db *sqlx.DB, table string) (bool, error) {
// 	query := fmt.Sprintf(`
//         SELECT EXISTS (
//             SELECT FROM information_schema.tables 
//             WHERE  table_schema = 'public'
//             AND    table_name   = '%s'
//         );
//     `, table)
// 	var exists bool 
// 	err := db.QueryRow(query).Scan(&exists)
// 	if err != nil {
// 		return false, err
// 	}
// 	return exists, nil
// }

// func createTable(db *sqlx.DB) error{
// 	sqlLines, err := readSQLFile("create_db.sql")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_,err = db.Exec(sqlLines)
// 	return err
// }

// func readSQLFile(file string) (string, error) {
// 	data, err := os.ReadFile(file)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(data), err
// }
