package main

import "fmt"

func main() {

	fmt.Println("Welcome to my quiz game!")

	// Enter name
	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	// Enter Age
	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	// Check the age
	if age >= 10 {
		fmt.Println("You can play!")
	} else {
		fmt.Println("You cannot play!")
		return
	}

	fmt.Println("What is better, Porto or Benfica?")
	var answer string
	fmt.Scan(&answer)

	fmt.Println(answer)

}
