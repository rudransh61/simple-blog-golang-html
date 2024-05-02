package main

import (
	"html/template"
	"net/http"
	"strconv"
)

// Post struct represents a blog post
type Post struct {
	ID      int
	Title   string
	Content string
}

// In-memory store for blog posts
var posts []Post

// Handler for the homepage
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Render the homepage template
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, struct{ Posts []Post }{Posts: posts})
}

// Handler for individual post pages
func postHandler(w http.ResponseWriter, r *http.Request) {
	// Extract post ID from URL
	id, _ := strconv.Atoi(r.URL.Path[len("/post/"):])
	post := posts[id]

	// Render the post template
	tmpl := template.Must(template.ParseFiles("post.html"))
	tmpl.Execute(w, post)
}

// Handler for creating new posts
func createHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	r.ParseForm()
	title := r.Form.Get("title")
	content := r.Form.Get("content")

	// Create new post
	post := Post{
		ID:      len(posts),
		Title:   title,
		Content: content,
	}
	posts = append(posts, post)

	// Redirect to homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	// Define routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/post/", postHandler)
	http.HandleFunc("/create", createHandler)

	// Start server
	http.ListenAndServe(":8080", nil)
}
