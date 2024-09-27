package pages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

// Beğenilen gönderi yapısını tanımla
type LikePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func fetchLikePosts() ([]LikePost, error) {
	resp, err := http.Get("http://localhost:8000/api/v1/mylikes")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// API yanıtını kontrol et
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var posts []LikePost
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// Ana handler
func IndexLikeHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := fetchLikePosts()
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	// HTML şablonunu oluştur
	tmpl := template.Must(template.ParseFiles("templates/mylikespost.html"))

	// Şablonu verilerle render et
	data := struct {
		Posts []LikePost
	}{
		Posts: posts,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
