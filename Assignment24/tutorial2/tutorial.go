package main

import "fmt"

func main() {
	fmt.Println("--- TUTORIAL STEP 1 ---")
	var x int            // declaration of a variable including type.
	x = 3                // assigns the variable x to 3.
	fmt.Println("x:", x) // prints 3

	//declare an integer variable y with value 20 in a single statement then print it.
	var y int = 20
	fmt.Println("y:", y)

	//declare variable z with value 50 and print it, type is not mentioned.
	var z = 50
	fmt.Println("z:", z)

	//declare ultiple variables and assign in single line, i with an integer and j with a string
	var i, j = 100, "hello"
	fmt.Println("i and j:", i, j)

	//variables can be declared using := instead of var, however this cannot be reassigned
	a := 20
	fmt.Println("a assigned with a := 20", a)

	//** a := 50 will throw an error at this point **

	//Constant Variables
	fmt.Println("\n--- TUTORIAL STEP 2 ---")

	const b = 10
	fmt.Println("Constant Variable b:", b)
	//cannot reassign b to another value: b = 20 will throw an error.

	//LOOPS
	fmt.Println("\n--- TUTORIAL STEP 3 ---")
	fmt.Println("--For Loop--")
	var iterator int
	for iterator = 1; iterator <= 5; iterator++ {
		fmt.Println(iterator)
	}

	//if-else
	fmt.Println("--If-Else-")
	var test = 50
	if test < 10 {
		fmt.Println("test is less than 10")
	} else if test >= 10 && test <= 90 {
		fmt.Println("test is between 10 and 90")
	} else {
		fmt.Println("test is greater than 90")
	}

	//switch
	fmt.Println("--Switch--")
	c, d := 2, 1
	switch c + d {
	case 1:
		fmt.Println("Sum is 1")
	case 2:
		fmt.Println("Sum is 2")
	case 3:
		fmt.Println("Sum is 3")
	default:
		fmt.Println("Printing default")
	}

	//Arrays
	fmt.Println("--Arrays--")
	var array [3]string
	array[0] = "one"
	array[1] = "two"
	array[2] = "three"
	fmt.Println(array[1])
	fmt.Println(len(array))
	fmt.Println(array)

	fmt.Println("--Functions--")
	num1, num2 := 15, 10
	sum, diff := calc(num1, num2)
	fmt.Println("Sum:", sum)
	fmt.Println("Diff:", diff)

}

func calc(num1 int, num2 int) (int, int) {
	fmt.Println("Inside the Function:")
	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}
