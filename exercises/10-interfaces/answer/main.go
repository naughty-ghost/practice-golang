package main

import (
	"fmt"
	"math"
)

// インターフェース定義
type Shape interface {
	Area() float64
}

// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Shape インターフェースを受け取る関数
func PrintArea(s Shape) {
	fmt.Printf("面積: %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	PrintArea(c)
	PrintArea(r)
}
