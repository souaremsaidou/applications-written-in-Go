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

	//var arr[3] int
	//arr[0] = 1
	// or
	arr := [3]int{1, 2, 3}
	fmt.Println(arr) // fmt.Println(arr[1])

	/*
	 * Slices
	 * dynamic size
	 */

	// slice from built-in array
	slice := arr[:]
	arr[0] = -1
	slice[2] = 0
	fmt.Println(arr, slice)

	// implicit
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	// slice is not fixed size
	// std::vector in c++
	slice2 = append(slice2, 4, 42)

	// not include index 2
	s := slice2[1:4]
	fmt.Println(slice2, s)

	/*
	 * Maps
	 * key, value
	 */

	m := map[string]int{"key": 0}
	// var m map[string]int //panic: assignment to entry in nil map
	m["key"] = 1 // map[key:1]
	//delete(m, "key")
	fmt.Println(m) // map[]

	/*
	 * Structs
	 * fixed number of fields
	 */
	type MyStruct struct {
		ID        int
		FirstName string
	}

	var me MyStruct
	//me.FirstName  = "Saidou"

	you := MyStruct{
		ID:        1,
		FirstName: "Diamilatou",
	}

	fmt.Println(me, you)

}
