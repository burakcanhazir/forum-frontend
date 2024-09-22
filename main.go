package main

import (
	"fmt"
	"log"
	"net/http"

	"forumfront/pages"
)

func main() {
	http.Handle("/", http.HandlerFunc(pages.IndexHandler))

	// Static dosyaları sunmak için
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("static/style"))))

	// Sunucuyu başlat
	fmt.Println("Server starting at http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
