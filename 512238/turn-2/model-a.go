package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type Point struct {
	X float64
	Y float64
}

func Distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Hypot(dx, dy)
}

func GetDistance(w http.ResponseWriter, r *http.Request) {
	p1X, err := fmt.ParseFloat(r.URL.Query().Get("x1"), 64)
	p1Y, err := fmt.ParseFloat(r.URL.Query().Get("y1"), 64)
	p2X, err := fmt.ParseFloat(r.URL.Query().Get("x2"), 64)
	p2Y, err := fmt.ParseFloat(r.URL.Query().Get("y2"), 64)

	if err != nil {
		http.Error(w, "Invalid input: Could not parse coordinates.", http.StatusBadRequest)
		return
	}

	p1 := Point{X: p1X, Y: p1Y}
	p2 := Point{X: p2X, Y: p2Y}

	distance := Distance(p1, p2)

	response := map[string]float64{"distance": distance}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error: Could not marshal JSON.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/distance", GetDistance)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
