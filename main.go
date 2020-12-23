package main

import (
	"fmt"
)

const ip = 3.14 // accessible to the entire pkg

// const blocks
const (
	A = iota + 6 // 0 + 6
	B            // 1 + 7
)

func main() {
	/*
	 * Primitives
	 */

	// var i int
	// i = 42
	var i int = 42
	fmt.Println(i)

	// implicit init
	firstName := "Saidou"
	fmt.Println(firstName)

	/*
	 * Pointer data type
	 */

	var p *int
	fmt.Println(p) // <nil>
	// panic: runtime error:
	// *p = 42 // invalid memory address or nil pointer dereference
	p = &i
	fmt.Println(*p)
	// invalid operation: p + 1
	//fmt.Println(p + 1)

	/*
	 * const
	 */

	//const
	const pi = 3.1415
	// pi = 1.2 //cannot assign to pi
	fmt.Println(pi)
	const c int = 3
	fmt.Println(c + 3)
	// fmt.Println(c + 1.2) // constant 1.2 truncated to integer
	//fmt.Println(float32(c) + 1.2) // implicit

	fmt.Println(c + 3)

	// const outside main
	fmt.Println(ip)

	// const blocks
	fmt.Println(A, B)

	/*
	 * Arrays
	 * fixed size 
	 * similar type
	 */
	
	// long syn
	var arr[3] int
	arr[0] = 1
	fmt.Println(arr) // fmt.Println(arr[1])
 
	
	
	

	/*
	 * Slices
	 * dynamic size
	 */

	/*
	 * Structs
	 * fixed number of fields
	 */

}
