package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func substract(x, y int) int { //Both param share type we can define type at the end
	return x - y
}

func swap(x, y string) (string, string) { //Multiple return
	return y, x
}

func split(sum int) (x, y int) { //Named results
	var i, j int = 4, 9
	k := 2 //Skip var inside fnc using := Outside a function, every construct begins
	//with a keyword (var, func, and so on) and the := construct is not available.
	const l = 1 //Constants are defined using const and cannot use :=
	x = sum*i/j + k*l
	y = sum - x - k*l
	return
}

func main() {
	fmt.Println(add(42, 13))

	fmt.Println(substract(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
