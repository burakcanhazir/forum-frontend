package main

import (
	"fmt"
	"log"
	"net/http"

	"forumfrontend/pages"
)

func main() {
	// createpost.html dosyasını sunmak için
	http.HandleFunc("/createpost", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/createpost.html")
	})
	http.HandleFunc("/getpost", pages.SelectedfetchPosts)

	// Statik dosyaları sunmak için
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("static/style"))))

	// Anasayfa için handler
	http.HandleFunc("/", pages.IndexHandler)

	// Sunucuyu başlat
	fmt.Println("Server starting at http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
