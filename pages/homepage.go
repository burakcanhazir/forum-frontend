package pages

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Post struct {
	PostID  string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// API'den postları çek
func fetchPosts() ([]Post, error) {
	resp, err := http.Get("http://localhost:8000/api/v1/homepage")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []Post
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
