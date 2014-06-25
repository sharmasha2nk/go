package main

import (
	"fmt"
)

func main() {
	var a, b = "0", "1"
	fmt.Println("Before swapping: \na=" + a + "\nb=" + b)
	a, b = b, a //swapping
	fmt.Println("After swapping: \na=" + a + "\nb=" + b)
}

/*
Output:

Before swapping:
a=0
b=1
After swapping:
a=1
b=0
*/
