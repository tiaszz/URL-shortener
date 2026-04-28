package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

func main() {
	// urls := make(map[string]string)
	// http.HandleFunc("/", helloHandler)
	// http.HandleFunc("POST /shorten", shortenHandler(urls))
	// http.HandleFunc("GET /{code}", redirect(urls))
	//
	// log.Fatal(http.ListenAndServe(":8080", nil))

	db, err := CreateDatabase()
	if err != nil {
		log.Fatal("Failed to create/open database: ", err)
	}
	defer db.Close()

	shortUrl := "4324dw"
	longUrl := "google.com"

	err = InsertData(db, shortUrl, longUrl)
	if err != nil {
		log.Fatal("Failed to insert data: ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func shortenHandler(urls map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}

		longURL := r.FormValue("url")

		shortened, err := generateMap(longURL, urls)
		if err != nil {
			return
		}

		fmt.Fprintf(w, "you shortened url is: localhost:8080/%s", shortened)
		fmt.Println(urls)
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

func generateMap(originalUrl string, urls map[string]string) (string, error) {
	shortenUrl := randomCode()
	_, exists := urls[shortenUrl]
	if exists {
		return "", errors.New("The shortened version for this is already created")
	}

	urls[shortenUrl] = originalUrl
	return shortenUrl, nil
}

func redirect(urls map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")
		value, exists := urls[code]
		if !exists {
			http.Error(w, "You didn't shortened your code", http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, value, http.StatusMovedPermanently)
	}
}
