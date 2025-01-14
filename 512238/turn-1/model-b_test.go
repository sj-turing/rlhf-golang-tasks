package distance

import (
	"math"
	"testing"
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

func TestDistance(t *testing.T) {
	testCases := []struct {
		name  string
		p1    Point
		p2    Point
		want  float64
		error error
	}{
		{
			name: "Origin to Origin",
			p1:   Point{X: 0, Y: 0},
			p2:   Point{X: 0, Y: 0},
			want: 0,
		},
		{
			name: "Positive X and Y",
			p1:   Point{X: 3, Y: 4},
			p2:   Point{X: 0, Y: 0},
			want: 5,
		},
		{
			name: "Negative X and Y",
			p1:   Point{X: -3, Y: -4},
			p2:   Point{X: 0, Y: 0},
			want: 5,
		},
		{
			name: "Positive X and Negative Y",
			p1:   Point{X: 3, Y: -4},
			p2:   Point{X: 0, Y: 0},
			want: 5,
		},
		{
			name: "Negative X and Positive Y",
			p1:   Point{X: -3, Y: 4},
			p2:   Point{X: 0, Y: 0},
			want: 5,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Distance(tt.p1, tt.p2)
			if err != nil {
				t.Errorf("Distance() error = %v, want nil", err)
				return
			}
			if got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
