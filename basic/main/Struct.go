package main

import "fmt"

type Vertex struct {
	x, y int
	z    string
}

var (
	p = Vertex{1, 2, "3"}  // has type Vertex
	q = &Vertex{1, 2, "3"} // has type *Vertex
	r = Vertex{x: 1}       // Y:0 is implicit
	s = Vertex{}           // X:0 and Y:0
)

func main() {
	fmt.Println(Vertex{1, 2, "3"})

	v := Vertex{}
	v.x = 4
	fmt.Println(v.x)

	fmt.Println(p, q, r, s)

	v1 := new(Vertex) //The expression new(T) allocates a zeroed T value and returns a pointer to it.
	//var t *T = new(T) or t := new(T)
	fmt.Println(v1)
	v1.x, v1.y, v1.z = 11, 9, "7"
	fmt.Println(v1)
}

/*
Output:

{1 2 3}
4
{1 2 3} &{1 2 3} {1 0 } {0 0 }
&{0 0 }
&{11 9 7}
*/
