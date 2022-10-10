package triangle

import "testing"

// First Test
func TestPerimeterTriangle(t *testing.T) {
	t.Run("Triangle perimeter 3, 4 and 6", func(t *testing.T) {
		got := Perimeter(Triangle{3, 4, 6, 0, 0})
		if got != 13 {
			t.Error("It's not good !!!")
		}
	})
}

func TestAreaTriangle(t *testing.T) {
	t.Run("Triangle area 3 and 4", func(t *testing.T) {
		got := Area(Triangle{0, 0, 0, 3, 4})
		if got != 6 {
			t.Error("It's not good !!!")
		}
	})
}

// Second Test
func TestTriangle_Perimeter(t *testing.T) {
	t.Run("Perimeter 3, 4 and 6", func(t *testing.T) {
		s := Triangle{
			side1: 3,
			side2: 4,
			side3: 6,
		}
		got := s.Perimeter()
		if got != 13 {
			t.Error("Perimeter 3, 4 and 6 is not 13")
		}
	})
}

func TestTriangle_Area(t *testing.T) {
	t.Run("Area  3 and 4", func(t *testing.T) {
		s := Triangle{
			height: 3,
			base:   4,
		}
		got := s.Area()
		if got != 6 {
			t.Error("Area 3 and 4 is not 6")
		}
	})
}