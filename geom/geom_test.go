package geom

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	cases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Perimeter 5 ", Pentagone{5, 3}, 25},
		{"Perimeter 2.5 ", Circle{2.5}, 15.707963267948966},
		{"Perimeter 4 and 6", Rectangle{4, 6}, 20},
		{"Perimeter 7 and 9", Rectangle{7, 9}, 32},
		{"Perimeter 4", Hexagone{4}, 24},
		{"Perimeter 12", Hexagone{12}, 72},
		{"Perimeter 3, 4 and 6", Triangle{3, 4, 6, 0, 0}, 13},
		{"perimeter 5 and 5", Star{5}, 50},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.shape.Perimeter()
			if got != c.want {
				t.Fatalf("got %f, want %f", got, c.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	cases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Area 5 and 3", Pentagone{5, 3}, 37.5},
		{"Area 4 and 6", Rectangle{4, 6}, 24},
		{"Area 7 and 9", Rectangle{7, 9}, 63},
		{"Area 5.1", Circle{5.1}, 81.71282},
		{"Area 4", Hexagone{4}, 41.569219381653056},
		{"Area 12", Hexagone{12}, 374.1229744348775},
		{"Area 3 and 4", Triangle{0, 0, 0, 3, 4}, 6},
		{"Area 5 and 5", Star{5}, 250},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.shape.Area()
			if fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", c.want) {
				t.Fatalf("got %.5f, want %.5f", got, c.want)
			}
		})
	}
}
