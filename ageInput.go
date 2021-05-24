// Created by: haokai
// Created on: May 2021
// This program displays, "Age-and-Movie"

package main

import (
	"fmt"
)

func main() {

	// This function does addition
	var age int

	// input
	fmt.Println("This program is about age limit of movie .")
	fmt.Println()
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	// detect
	if age >= 17 {
		// output of right answer
		fmt.Println("You can see a R movie alone.")
	} else if age >= 13 {
		// output of wrong answer
		fmt.Println("You can see a PG-13 movie alone.")
	} else if age >= 5 {
		// output of wrong answer
		fmt.Println("You can see a G or PG movie alone.")
	} else {
		// output of wrong answer
		fmt.Println("Uh, you are too young for most things.")
	}
	fmt.Println("Thanks for verifying  your age!")
	fmt.Println("\nDone.")
}
