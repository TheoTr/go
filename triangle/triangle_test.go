package triangle

import "testing"

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
