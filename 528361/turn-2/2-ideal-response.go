package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Blog struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

type BlogByDate []*Blog

var (
	blogs      = []*Blog{}
	nextBlogID int
)

func (b BlogByDate) Len() int {
	return len(b)
}

func (b BlogByDate) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b BlogByDate) Less(i, j int) bool {
	return b[i].Created.Before(b[j].Created)
}

// FilterByYear filters blogs by a given year
func (b BlogByDate) FilterByYear(year int) BlogByDate {
	var filteredBlogs BlogByDate
	for _, blog := range b {
		if blog.Created.Year() == year {
			filteredBlogs = append(filteredBlogs, blog)
		}
	}
	return filteredBlogs
}

// FilterByYearMonth filters blogs by a given year and month
func (b BlogByDate) FilterByYearMonth(year, month int) BlogByDate {
	var filteredBlogs BlogByDate
	for _, blog := range b {
		if blog.Created.Year() == year && blog.Created.Month() == time.Month(month) {
			filteredBlogs = append(filteredBlogs, blog)
		}
	}
	return filteredBlogs
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

func createBlog(w http.ResponseWriter, r *http.Request) {
	var blog Blog
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	blog.ID = nextBlogID
	nextBlogID++
	blog.Created = time.Now()
	blogs = append(blogs, &blog)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&blog)
}

func getBlogsByDate(w http.ResponseWriter, r *http.Request) {
	year, err := strconv.Atoi(mux.Vars(r)["year"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	month, err := strconv.Atoi(mux.Vars(r)["month"])
	if err != nil {
		filteredBlogs := BlogByDate(blogs).FilterByYear(year)
		if len(filteredBlogs) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(filteredBlogs)
		return
	}
	day, err := strconv.Atoi(mux.Vars(r)["day"])
	if err != nil {
		filteredBlogs := BlogByDate(blogs).FilterByYearMonth(year, month)
		if len(filteredBlogs) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(filteredBlogs)
		return
	}

	filteredBlogs := BlogByDate(blogs).FilterByYearMonthDay(year, month, day)
	if len(filteredBlogs) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredBlogs)
}

func getBlogByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	blog := BlogByDate(blogs).GetByID(id)
	if blog == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/blogs", createBlog).Methods("POST")
	r.HandleFunc("/blogs/{year}/{month}/{day}", getBlogsByDate).Methods("GET")
	r.HandleFunc("/blogs/{year}/{month}", getBlogsByDate).Methods("GET")
	r.HandleFunc("/blogs/{year}", getBlogsByDate).Methods("GET")
	r.HandleFunc("/blog", getBlogByID).Methods("GET")

	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
