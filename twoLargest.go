package main

import "fmt"

func main() {
	var largest, secondLargest int = -99999, -99999

	for count := 1; count <= 10; count++ {
		var number int
		fmt.Printf("Enter number %d: ", count)
		fmt.Scanln(&number)

		if number > largest {
			secondLargest = largest
			largest = number
		} else if number > secondLargest {
			secondLargest = number
		}
	}

	fmt.Printf("\nThe two largest numbers are: %d and %d\n", largest, secondLargest)
}
