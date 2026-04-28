package main

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

func CreateDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./database/links.db")
	if err != nil {
		return nil, err
	}

	return db, nil

}

func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE link(
		shortened VARCHAR(6) NOT NULL,
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
