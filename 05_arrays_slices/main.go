package main

import "fmt"

func main()  {

	//var fruitArr [2] string

	// fruitArr[0] = "orange"
	// fruitArr[1] = "apple"

	//declare and assign
	fruitArr := [2]string {"apple", "orange"}
	
	fruitSlices := []string{"apple", "orange", "grape"}
	//fmt.Println(fruitArr)
	fmt.Println(fruitSlices[1:2])

	for i := 0; i < len(fruitArr); i++ {
		fmt.Println(fruitArr[i])
	}
}