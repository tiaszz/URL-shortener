package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

func main() {
	urls := make(map[string]string)
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/shorten", getUrlHandler(urls))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func getUrlHandler(urls map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}

		longURL := r.FormValue("url")

		err := generateMap(longURL, urls)
		if err != nil {
			return
		}

		fmt.Fprintf(w, "you shortened url is: %s", urls[longURL])
		fmt.Println(urls)
	}
}

func randomCode() string {
	random := "localhost:8080/"
	useForRandom := "abcdefghijklmnoqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for range 6 {
		random += string(useForRandom[rand.IntN(len(useForRandom))])
	}

	return random
}

func generateMap(originalUrl string, urls map[string]string) error {
	shortenUrl := randomCode()
	_, exists := urls[originalUrl]
	if exists {
		return errors.New("The shortened version for this is already created")
	}

	urls[originalUrl] = shortenUrl
	return nil
}
