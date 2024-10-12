package main

import (
	"fmt"
	"log"
	"net/http"

	"forumfront/pages"
)

func main() {
	// Anasayfa için handler
	http.Handle("/", http.HandlerFunc(pages.IndexHandler))

	// createpost.html dosyasını sunmak için
	http.HandleFunc("/createpost", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/createpost.html")
	})

	// Static dosyaları sunmak için
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("static/style"))))

	// Sunucuyu başlat
	fmt.Println("Server starting at http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
