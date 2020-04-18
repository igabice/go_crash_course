package main

import "fmt"

func greetings(name string) string {
	return "hello "+name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main()  {
	fmt.Println(greetings("israel"))
	fmt.Println(getSum(4, 11))
}