package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

func (p Person) greet() string  {
	return "hello " +p.firstName+ strconv.Itoa(p.age)
	
}
func (p *Person) hasBirthday()  {
	p.age++
}

func main()  {
	
var person1 = Person{"smith", "samantha", "boston", "f", 5}
person1.hasBirthday()	
fmt.Println(person1.greet()) 
	
}  