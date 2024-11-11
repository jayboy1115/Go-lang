package main

import "fmt"

func main() {
	for {
		fmt.Print("Enter 1 for integer input or 2 for float input (or -1 to quit): ")
		var choice int
		fmt.Scanln(&choice)

		if choice == -1 {
			break
		}

		for choice != 1 && choice != 2 {
			fmt.Println("Invalid choice. Please enter 1 or 2.")
			fmt.Print("Enter 1 for integer input or 2 for float input: ")
			fmt.Scanln(&choice)
		}

		if choice == 1 {
			var integer int
			fmt.Print("Enter an integer: ")
			fmt.Scanln(&integer)
			fmt.Printf("You entered %d\n", integer)
		} else if choice == 2 {
			var floatingPoint float64
			fmt.Print("Enter a floating-point number: ")
			fmt.Scanln(&floatingPoint)
			fmt.Printf("You entered %f\n", floatingPoint)
		}
	}
}
