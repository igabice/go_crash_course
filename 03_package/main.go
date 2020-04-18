package main

import (
	"fmt"
	"math"
	"github.com/igabice/go_crash_source/03_package/strutil"
)

func main()  {
	fmt.Println(math.Floor(1.32))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))

	fmt.Println(strutil.Reverse("Hello"))
} 