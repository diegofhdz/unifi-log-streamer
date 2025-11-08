package main

import "fmt"

func main() {
	num := 10
	var ptr *int = &num // Declare a pointer to an integer

	fmt.Println("Value of num:", num)             // Output: Value of num: 10
	fmt.Println("Address of num:", &num)          // Output: Address of num: 0xc000014088 (example address)
	fmt.Println("Value of ptr:", ptr)             // Output: Value of ptr: 0xc000014088 (example address)
	fmt.Println("Value pointed to by ptr:", *ptr) // Output: Value pointed to by ptr: 10

	*ptr = 20 // Change the value at the address pointed to by 'ptr'

	fmt.Println("New value of num:", num) // Output: New value of num: 20

	num = 30
	fmt.Println("Num", num)
}
