package main

import "fmt"

func main()  {
	
	fmt.Println(" ")

	// emails := make(map[string]string)

	// emails["bob"] = "bob@gmail.com"
	// emails["sharon"] = "sharon@gmail.com"
	// emails["mike"] = "mike@gmail.com"
	
	emails := map[string]string{"bob":"bob@gmail.com", "sharon":"sharon@gmail.com",}
	//delete
	delete(emails, "bob")
	fmt.Println(emails)
	fmt.Println(emails["bob"])
}