package main

import 
(
	"fmt" 
	"math"
)

type Circle struct {
	x, y, radius float64
}

type Shape interface {
	area()
}

type Rectangle struct{
	width, height float64
}  

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
 
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main()  {
	circle := Circle{0,0,9}
	rectangle := Rectangle{ 10,5}


	fmt.Println(getArea(circle));
	fmt.Println(getArea(rectangle));
	
} 