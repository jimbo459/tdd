package shapes

import(
"math"
)

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
  Width, Height float64
}

type Shape interface {
  Area() float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func(c Circle) Area() float64 {
  return  math.Pi * c.Radius * c.Radius 
}

func(r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func(t Triangle) Area() float64 {
  return (t.Width / 2) * t.Height 
}
