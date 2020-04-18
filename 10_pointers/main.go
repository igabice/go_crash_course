package main

import "fmt"

func main()  {
	
	a := 5
	b := &a
	fmt.Println(*b);
	//pass by value
	//faster to call address where var is saved
	
}