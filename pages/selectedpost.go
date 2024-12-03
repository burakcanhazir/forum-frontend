package pages

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type SelectedPost struct {
	ID        string        `json:"id"` // Post ID'si
	Title     string        `json:"title"`
	UserID    string        `json:"user_id"` // Postu oluşturan kullanıcının ID'si
	Content   string        `json:"content"` // Post içeriği
	Category  []string      `json:"category"`
	CreatedAt string        `json:"created_at"` // Postun oluşturulma tarihi
	Likes     sql.NullInt64 `json:"likes"`      // Nullable integer type
}

func SelectedfetchPosts(w http.ResponseWriter, r *http.Request) {
	log.Printf("BURASI selected")
	// URL'den 'id' parametresini al
	postID := r.URL.Query().Get("id")
	if postID == "" {
		http.Error(w, "Missing 'id' parameter in URL", http.StatusBadRequest)
		return
	}

	// Ortam değişkeninden backend URL'sini al
	baseURL := os.Getenv("BACKEND_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8000" // Varsayılan olarak localhost
	}

	// URL'yi oluştur
	url := fmt.Sprintf("%s/api/v1/getpost/%s", baseURL, postID)
	fmt.Println("Fetching from URL:", url)

	// API'ye GET isteği yap
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching post: %v", err)
		http.Error(w, "Failed to fetch post", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// HTTP yanıtının durum kodunu kontrol et
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Received status code %d from backend", resp.StatusCode)
		http.Error(w, "Failed to fetch post: Non-OK status code", http.StatusInternalServerError)
		return
	}

	// JSON verisini çözümlemek için bir struct oluştur
	var post SelectedPost
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}
	log.Println("Response Body:", string(body))

	// JSON'u çözümle ve `post` struct'ına aktar
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		http.Error(w, "Failed to unmarshal JSON", http.StatusInternalServerError)
		return
	}

	// HTML şablonunu verilerle render et
	tmpl := template.Must(template.ParseFiles("templates/userpost.html"))
	data := struct {
		Post SelectedPost
	}{
		Post: post,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
