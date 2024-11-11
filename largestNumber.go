package main

import "fmt"

func main() {
	var counter int = 0
	var number, largest int = 0, -99999

	for counter < 5 {
		fmt.Print("Enter number: ")
		fmt.Scanln(&number)

		if number > largest {
			largest = number
		}

		counter++
	}

	fmt.Printf("Largest number is: %d\n", largest)
}
