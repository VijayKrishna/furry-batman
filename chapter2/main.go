package main

import "fmt"

// this is a comment.
/* so is this*/

func main() {
	var message string = "Hello my name is Vijay"
	fmt.Println(message)
	fmt.Println(len("Hello my name is Vijay"))
	fmt.Println("Hello"[4])
	fmt.Println("Hello" + " World")

	fmt.Println(true && true)
	fmt.Println(false && true)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)

	var x string = "first"
	var y string = "yello"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
	fmt.Println(x + " but first")

	fmt.Println("hello" == "hello")
	fmt.Println(x == y)

	var z = "vijay"
	fmt.Println(z)

	zz := "yo yo"
	fmt.Println(zz)

	puppyname := "pepper"
	mastername := "vijay"

	fmt.Println(mastername + ", as a kid, had a puppy named " + puppyname)
}