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
	  areaTests := []struct{
      shape Shape
      want float64
    }{
      {Rectangle{12,6}, 72.0},
      {Circle{10}, 314.1592653589793},
      {Triangle{12,6}, 36.0},
    }
    for _, tt := range areaTests {
      got := tt.shape.Area()
      if got != tt.want {
        t.Errorf("%#v got %.2f, want %.2f",tt.shape, got, tt.want)
      }
    }
}
