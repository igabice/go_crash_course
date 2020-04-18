package main

import "fmt"

func main()  {

	ids := []int{22,11,33,41,5,66,77,44}

	for _, id := range ids {
		fmt.Printf(" ID: %d\n", id);
	}
	
	sum := 0
	for _, id := range ids {
		sum += id
	}
	//fmt.Println(sum)
	
	emails := map[string]string{"bob":"bob@gmail.com", "sharon":"sharon@gmail.com",}
	
	for k,v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	
	for k := range emails {
		fmt.Printf("Name: "+ k)
	}
}