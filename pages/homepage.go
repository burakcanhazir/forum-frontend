package pages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Post struct {
	PostID   string   `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category []string `json:"category"`
}

// API'den postları çek
func fetchPosts() ([]Post, error) {
	// Ortam değişkeninden backend URL'sini al
	baseURL := os.Getenv("BACKEND_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8000" // Varsayılan olarak localhost
	}

	// URL'yi oluştur
	url := fmt.Sprintf("%s/api/v1/homepage", baseURL)
	log.Println("Fetching from URL:", url)

	// API'ye GET isteği yap
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		return nil, fmt.Errorf("failed to fetch posts: %w", err)
	}
	defer resp.Body.Close()

	// HTTP yanıtının durum kodunu kontrol et
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Received status code %d from backend", resp.StatusCode)
		return nil, fmt.Errorf("received status code %d from backend", resp.StatusCode)
	}

	// JSON verisini çözümlemek için bir struct oluştur
	var posts []Post
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// JSON'u çözümle ve `posts` dilimlerine aktar
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return posts, nil
}

// Ana handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := fetchPosts()
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	// HTML şablonunu oluştur
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Şablonu verilerle render et
	data := struct {
		Posts []Post
	}{
		Posts: posts,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
