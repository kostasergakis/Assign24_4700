package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

//Variables:
var (
	c, java, python bool
	toBe            bool       = false
	maxInt          uint64     = 1<<64 - 1
	z               complex128 = cmplx.Sqrt(-1 + 12i)
)

func main() {
	fmt.Println("Hello there! *Kenobi voice*")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println("\n===Calling the function \"add\".===")
	fmt.Println(add(42, 13))
	fmt.Println("\n===Printing the variables.===")
	var i, j int = 1, 2
	var c, java, python = true, "no!", false
	fmt.Println(i, j, c, python, java)
	fmt.Println("\n===Printing the complex variables.===")
	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println("\n===Printing the for loop tests===")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println("Second Loop (for keyword is used for while loop.")
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)
	fmt.Println("\n===Printing the if statements===")
	if c {
		fmt.Println("This was true!")
	} else {
		fmt.Println("This was false!")
	}

	fmt.Println("\n===Printing the arrays(fixed size)===")
	var a [10]int
	fmt.Println(a)
	fmt.Println("\n===Printing the slices(dynamic size)===")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println("\n===Printing the slices(dynamic size) created with make keyword===")
	b := make([]int, 5)
	printSlice("b", b)
	d := make([]int, 0, 5)
	printSlice("d", d)
	e := b[:2]
	printSlice("e", e)

}

//Functions:

func add(x, y int) int {
	return x + y
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
