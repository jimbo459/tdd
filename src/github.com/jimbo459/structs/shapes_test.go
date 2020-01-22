package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Rectangles", func(t *testing.T) {
		got := Perimeter(Rectangle{Width: 10, Height: 10})
		want := 40.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

}

func TestArea(t *testing.T) {
	t.Run("Rectangles", func(t *testing.T) {
		got := Area(Rectangle{Width: 10, Height: 10})
		want := 100.00

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		got := Area(circle)
		want := 314.15

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}
