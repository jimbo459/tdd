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
	checkArea := func(t *testing.T, shape Shape, want float64){
    got := shape.Area()
    if got != want {
      t.Errorf("got %.2f, want %.2f", got, want)
    }
  }

  t.Run("Rectangles", func(t *testing.T) {
    rectangle := Rectangle{10,10}
    checkArea(t, rectangle, 100)
	})
	t.Run("Circles", func(t *testing.T) {
    circle := Circle{10}
    want := 314.1592653589793
    checkArea(t, circle, want)
	})
}
