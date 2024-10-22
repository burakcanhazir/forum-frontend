package pages

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	// Belirtilen URL'ye istek yap, ID'yi kullanarak
	url := fmt.Sprintf("http://localhost:8000/api/v1/getpost/%s", postID)
	fmt.Println("Fetching from URL:", url)

	// API'ye GET isteği yap
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch post", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// JSON verisini çözümlemek için bir struct oluştur
	var post SelectedPost
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
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
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
