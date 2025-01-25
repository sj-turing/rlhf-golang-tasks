package main

import (
	"fmt"
	"time"
)

// Blog struct to hold individual blog entries
type Blog struct {
	ID      int
	Title   string
	Content string
	Date    time.Time
}

// BlogsYear struct to hold blogs for a specific year
type BlogsYear struct {
	Year  int
	Blogs []Blog
}

// BlogsMonth struct to hold blogs for a specific month within a year
type BlogsMonth struct {
	Month int
	Blogs []Blog
}

// BlogsDay struct to hold blogs for a specific day within a month
type BlogsDay struct {
	Day   int
	Blogs []Blog
}

type BlogPlatform struct {
	blogs []Blog
}

// NewBlogPlatform initializes and returns a new BlogPlatform
func NewBlogPlatform() *BlogPlatform {
	return &BlogPlatform{
		blogs: make([]Blog, 0),
	}
}

// AddBlog adds a new blog to the platform
func (bp *BlogPlatform) AddBlog(id int, title, content string, date time.Time) {
	bp.blogs = append(bp.blogs, Blog{ID: id, Title: title, Content: content, Date: date})
}

// GetBlogsByYear retrieves all blogs for a specific year
func (bp *BlogPlatform) GetBlogsByYear(year int) []Blog {
	var result []Blog
	for _, blog := range bp.blogs {
		if blog.Date.Year() == year {
			result = append(result, blog)
		}
	}
	return result
}

// GetBlogsByMonth retrieves all blogs for a specific year and month
func (bp *BlogPlatform) GetBlogsByMonth(year, month int) []Blog {
	var result []Blog
	for _, blog := range bp.blogs {
		if blog.Date.Year() == year && blog.Date.Month() == time.Month(month) {
			result = append(result, blog)
		}
	}
	return result
}

// GetBlogsByDay retrieves all blogs for a specific year, month, and day
func (bp *BlogPlatform) GetBlogsByDay(year, month, day int) []Blog {
	var result []Blog
	for _, blog := range bp.blogs {
		if blog.Date.Year() == year && blog.Date.Month() == time.Month(month) && blog.Date.Day() == day {
			result = append(result, blog)
		}
	}
	return result
}

// GetBlogByID retrieves a blog by its ID
func (bp *BlogPlatform) GetBlogByID(id int) *Blog {
	for _, blog := range bp.blogs {
		if blog.ID == id {
			return &blog
		}
	}
	return nil
}

func main() {
	// Example usage
	bp := NewBlogPlatform()
	bp.AddBlog(1, "First Blog", "Content for the first blog", time.Date(2023, 10, 21, 0, 0, 0, 0, time.UTC))
	bp.AddBlog(2, "Second Blog", "Content for the second blog", time.Date(2023, 10, 22, 0, 0, 0, 0, time.UTC))
	bp.AddBlog(3, "Third Blog", "Content for the third blog", time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC))

	blogsInOct := bp.GetBlogsByMonth(2023, 10)
	for _, blog := range blogsInOct {
		fmt.Println(blog.Title)
	}

	blogByID := bp.GetBlogByID(2)
	if blogByID != nil {
		fmt.Println("Blog found by ID:", blogByID.Title)
	}
}
