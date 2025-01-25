package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Blog represents a single blog post
type Blog struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

// BlogByDate is a slice of blog pointers organized by date
type BlogByDate []*Blog

var (
	blogs      BlogByDate
	nextBlogID = 1
)

func main() {
	http.HandleFunc("/create", createBlogHandler)
	http.HandleFunc("/blog", getBlogByIDHandler)
	http.HandleFunc("/year/month/day", getBlogsByDateHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// createBlogHandler handles the creation of a new blog
func createBlogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var blog Blog
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&blog); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	blog.ID = nextBlogID
	blog.Created = time.Now()
	nextBlogID++

	blogs = append(blogs, &blog)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(blog)
}

// getBlogByIDHandler handles retrieval of a blog by its ID
func getBlogByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Bad request: missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad request: invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	blog := blogs.GetByID(id)
	if blog == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(blog)
}

// getBlogsByDateHandler handles retrieval of blogs by date
func getBlogsByDateHandler(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	monthStr := r.URL.Query().Get("month")
	dayStr := r.URL.Query().Get("day")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Bad request: invalid 'year' parameter", http.StatusBadRequest)
		return
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		http.Error(w, "Bad request: invalid 'month' parameter", http.StatusBadRequest)
		return
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		http.Error(w, "Bad request: invalid 'day' parameter", http.StatusBadRequest)
		return
	}

	blogsByDate := blogs.FilterByYearMonthDay(year, month, day)
	json.NewEncoder(w).Encode(blogsByDate)
}

// FilterByYearMonthDay filters blogs by a given year, month, and day
func (b BlogByDate) FilterByYearMonthDay(year, month, day int) BlogByDate {
	var filteredBlogs BlogByDate
	for _, blog := range b {
		if blog.Created.Year() == year && blog.Created.Month() == time.Month(month) && blog.Created.Day() == day {
			filteredBlogs = append(filteredBlogs, blog)
		}
	}
	return filteredBlogs
}

// GetByID retrieves a blog by its ID
func (b BlogByDate) GetByID(id int) *Blog {
	for _, blog := range b {
		if blog.ID == id {
			return blog
		}
	}
	return nil
}
