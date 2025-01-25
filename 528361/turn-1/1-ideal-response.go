package main

import (
	"fmt"
	"time"
)

// Blog represents a single blog post
type Blog struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

// BlogByDate is a slice of blogs organized by date
type BlogByDate []*Blog

// ByYear implements sort.Interface for BlogByDate
func (b BlogByDate) Len() int           { return len(b) }
func (b BlogByDate) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b BlogByDate) Less(i, j int) bool { return b[i].Created.Before(b[j].Created) }

func main() {
	// Initialize blogs
	blogs := []*Blog{
		{ID: 1, Title: "Blog 1", Content: "Content of Blog 1", Created: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
		{ID: 2, Title: "Blog 2", Content: "Content of Blog 2", Created: time.Date(2023, 2, 3, 0, 0, 0, 0, time.UTC)},
		{ID: 3, Title: "Blog 3", Content: "Content of Blog 3", Created: time.Date(2023, 1, 25, 0, 0, 0, 0, time.UTC)},
		{ID: 4, Title: "Blog 4", Content: "Content of Blog 4", Created: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	// Organize blogs by date
	blogsByDate := BlogByDate(blogs)

	// Display blogs by date
	fmt.Println("Blogs by date:")
	for _, blog := range blogsByDate {
		fmt.Printf("%s - %s\n", blog.Created.Format("2006-01-02"), blog.Title)
	}

	// Retrieve blogs by year
	blogsByYear := blogsByDate.FilterByYear(2023)
	fmt.Println("\nBlogs by year 2023:")
	for _, blog := range blogsByYear {
		fmt.Printf("%s - %s\n", blog.Created.Format("2006-01-02"), blog.Title)
	}

	// Retrieve blogs by year-month
	blogsByYearMonth := blogsByDate.FilterByYearMonth(2023, 1)
	fmt.Println("\nBlogs by year-month 2023-01:")
	for _, blog := range blogsByYearMonth {
		fmt.Printf("%s - %s\n", blog.Created.Format("2006-01-02"), blog.Title)
	}

	// Retrieve blogs by year-month-day
	blogsByYearMonthDay := blogsByDate.FilterByYearMonthDay(2023, 1, 1)
	fmt.Println("\nBlogs by year-month-day 2023-01-01:")
	for _, blog := range blogsByYearMonthDay {
		fmt.Printf("%s - %s\n", blog.Created.Format("2006-01-02"), blog.Title)
	}

	// Retrieve blog by ID
	blogByID := blogsByDate.GetByID(3)
	if blogByID != nil {
		fmt.Println("\nBlog by ID 3:")
		fmt.Printf("%s - %s\n", blogByID.Created.Format("2006-01-02"), blogByID.Title)
	} else {
		fmt.Println("\nBlog by ID 3 not found.")
	}
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
