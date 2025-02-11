package main

import (
	"fmt"
)

type Person struct {
	name  string
	age   int
	isOld bool
}

func main() {
	/**
	* Working with pointers
	**/

	// var x int = 10
	// var xPtr *int = &x
	// Declaring regular variable
	x := 10
	// Declaring pointer variable
	xPtr := &x
	// Declaring pointer variable (reference address of x)
	y := &x

	// Print the value of x
	fmt.Println(x)
	// Print the address of x
	fmt.Println(xPtr)
	// Print the equality of the pointer y and address of x
	fmt.Println("y == &x //", y == &x)
	// Print the equality of the value of x and the (dereferenced) value of the pointer y
	fmt.Println("x == *y //", x == *y)

	/**
	* Working with structs
	**/

	// Method 1: Value type, empty initialization
	person1 := Person{}
	person1.age = 20
	person1.name = "James"
	person1.isOld = false

	// Print the value of the Person struct
	fmt.Println(person1)

	// -----------------------------------

	// Method 2: Value type, field initialization
	person2 := Person{name: "Billy", age: 45, isOld: true}
	// Print the value of the Person struct
	fmt.Println(person2)

	// -----------------------------------

	// Method 3: Pointer type initialization
	person3 := &Person{name: "Victor", age: 32, isOld: false}
	// Print the value of the Person struct
	fmt.Println(person3)

	// -----------------------------------

	// Test passing copy of person
	testStructModificationCopy(*person3)
	// Ensure that the original person is not modified
	fmt.Println(person3)

	// Test passing pointer of person
	testStructModificationPointer(person3)
	// Ensure that the original person is modified
	fmt.Println(person3)
}

func testStructModificationCopy(person Person) Person {
	person.age = 90
	return person
}

func testStructModificationPointer(person *Person) Person {
	person.age = 90
	return *person
}
