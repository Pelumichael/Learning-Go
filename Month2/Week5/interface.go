package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	name           string
	length, height float64
}

func (r *rectangle) area() float64 {
	return r.length * r.height
}

func (r *rectangle) perimeter() float64 {
	return 2*r.length + 2*r.height
}

type triangle struct {
	name    string
	a, b, c float64
}

func (t *triangle) area() float64 {
	return 0.5 * (t.a * t.b)
}

func (t *triangle) perimeter() float64 {
	return t.a + t.b + math.Sqrt((t.a*t.a)+(t.b*t.b))
}

func (t *triangle) String() string {
	return fmt.Sprintf("%s[sides: a=%.2f b=%.2f c=%.2f]", t.name, t.a, t.b, t.c)
}

// Using string Method
func shapeInfo(s Shape) string {
	return fmt.Sprintf("Area = %.2f, Perimeter = %.2f", s.area(), s.perimeter())
}

func main() {
	r := &rectangle{"First Rectangle", 4, 4}
	fmt.Println(r, "=>", shapeInfo(r))

	t := &triangle{
		name: "First Triangle",
		a:    1,
		b:    2,
		c:    3,
	}
	fmt.Println(t, "=>", shapeInfo(t))
}
