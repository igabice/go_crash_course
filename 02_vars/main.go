package main

import "fmt"

func main()  {
	/*
	string
	bool
	int
	int int8  int16 int32 int64
	uint uint8  uint16 uint32 uint64 uintptr
	byte - alias for uint8
	rune alias for int32
	float32 float64
	complex64 complex128 really large nos
	
	*/
	//var name  = "israel"
	
	name, email := "israel", "igabice@gmail.com"
	var age int32 = 37
	const iscool= true
	fmt.Println("hello "+name, age, email)
	fmt.Printf ("%T\n", iscool)
}