package main

import "math"

// pass with updateVal(&val)
func updateVal(x *int) {
	*x += 5
}

type Rectangle struct {
	leftTopX int
	leftTopY int
	height   int
	width    int
}

func (rect *Rectangle) rightPoints() (int, int) {

	return rect.width + rect.leftTopX, rect.height + rect.leftTopY

}

type Shape interface {
	area() float64
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func Less(s1, s2 Shape) Shape {
	if s1.area() < s2.area() {
		return s1
	}
	return s2
}

type Circle struct {
	radius float64
}

type Triangle struct {
	heightB float64
	base    float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (t Triangle) area() float64 {
	return t.heightB * t.base / 2
}
