package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Reactangle struct {
	width  float64
	height float64
}

func (r Reactangle) Area() float64 {
	return r.width * r.height
}

func (r Reactangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func CalculateArea(s Shape, val string) float64 {
	if val == "a" {
		return s.Area()
	} else if val == "p" {
		return s.Perimeter()
	} else {
		return 0
	}

}

func main() {
	/* Here we are type casting the Reactabngle structure into Shape interface and go only allows to do when the structure implements
	all the methods of the interface.
	*/
	shape1 := Shape(Reactangle{width: 10, height: 5})
	/* Here we are type casting the Circle structure into Shape interface and go only allows to do when the structure implements
	all the methods of the interface.
	*/
	fmt.Println(shape1.Area())
	shape2 := Shape(Circle{radius: 7})
	fmt.Println(shape2.Area())

	// Or You can do in below way
	fmt.Println("Area of Circle:", CalculateArea(Circle{radius: 7}, "a"))
	fmt.Println("Area of Rectangle:", CalculateArea(Reactangle{width: 10, height: 5}, "a"))
	fmt.Println("Perimeter of Circle:", CalculateArea(Circle{radius: 7}, "p"))
}
