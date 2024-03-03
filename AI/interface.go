package main

import (
	"fmt"
	"math"
)

type area interface {
	area()float64
}
type Circle struct{
	circumference float64
}
type Rect struct{
	long float64
	short float64
}
func(c Circle) area() float64{
	return c.circumference * c.circumference * math.Pi
}

func (r Rect) area() float64{
	return r.long*r.short
}
func main() {
	c1 := Circle{5.04}
	r1 := Rect{13.4,18.1}
	fmt.Println(area.area(c1))
	fmt.Println(area.area(r1))
	
}