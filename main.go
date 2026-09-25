package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	if target == "" {
		fmt.Fprintf(w, "Provide a ?url= parameter.")
		return
	}

	isSecure := os.Getenv("SECURE_MODE") == "true"

	if isSecure {
		
		parsedURL, err := url.Parse(target)
		if err != nil || parsedURL.IsAbs() {
			http.Error(w, "External redirects are forbidden!", http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, target, http.StatusFound)
	} else {
		
		http.Redirect(w, r, target, http.StatusFound)
	}
}

func main() {
	http.HandleFunc("/redirect", redirectHandler)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
