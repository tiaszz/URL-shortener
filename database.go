package main

import (
	"database/sql"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

// Function to create a database at ./database/links.db
func CreateDatabase(dbName string) (*sql.DB, error) {
	path := "./database/" + dbName
	os.Mkdir("./database", 0o755)

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Function to create a table link at the database
// The schema is:
// Link (
//
//	shortened varchar(6) PK
//	long_url varchar(255)
//	created_at CURRENT_TIMESTAMP
//
// )
func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS link(
		shortened VARCHAR(6) NOT NULL PRIMARY KEY,
		long_url VARCHAR(255) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// Function to inset values in the link table
func InsertData(db *sql.DB, shortUrl, longUrl string) error {
	query := `
	INSERT INTO link(shortened, long_url, created_at)
	VALUES(?, ?, ?);
	`

	_, err := db.Exec(query, shortUrl, longUrl, time.Now())
	if err != nil {
		return err
	}

	return nil
}

// Function to delete a value in the link table
func DeleteData(db *sql.DB, shortUrl string) error {
	query := `
	DELETE FROM link WHERE shortened=?;
	`

	_, err := db.Exec(query, shortUrl)
	if err != nil {
		return err
	}

	return nil
}

func SelectData(db *sql.DB, code string) (string, error) {
	query := `
		SELECT long_url
		FROM link
		WHERE shortened=?;
		`
	row := db.QueryRow(query, code)

	var value string
	err := row.Scan(&value)
	if err != nil {
		return "", err
	}

	return value, nil
}
