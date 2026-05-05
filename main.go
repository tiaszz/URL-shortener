package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

func main() {
	db, err := CreateDatabase("links.db")
	if err != nil {
		log.Fatal("Failed to create/open database: ", err)
	}
	defer db.Close()

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("POST /shorten", shortenHandler(db))
	http.HandleFunc("GET /{code}", redirect(db))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func shortenHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}

		longURL := r.FormValue("url")
		shortenUrl := randomCode()

		err := InsertData(db, shortenUrl, longURL)
		if err != nil {
			return
		}

		fmt.Fprintf(w, "you shortened url is: localhost:8080/%s", shortenUrl)
	}
}

func randomCode() string {
	random := ""
	useForRandom := "abcdefghijklmnoqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for range 6 {
		random += string(useForRandom[rand.IntN(len(useForRandom))])
	}

	return random
}

func redirect(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")
		value, err := SelectData(db, code)
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, value, http.StatusMovedPermanently)
	}
}
