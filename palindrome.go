package main

import "fmt"

func main() {
	for {
		fmt.Print("Enter a five-digit integer: ")
		var number int
		fmt.Scanln(&number)

		if number < 10000 || number > 99999 {
			fmt.Println("Error: Number must be five digits long.")
			continue
		}

		str := fmt.Sprintf("%d", number)

		if isPalindrome(str) {
			fmt.Printf("%d is a palindrome.\n", number)
			break
		} else {
			fmt.Printf("%d is not a palindrome.\n", number)
			break
		}
	}
}

func isPalindrome(s string) bool {
	for count := 0; count < len(s)/2; count++ {
		if s[count] != s[len(s)-1-count] {
			return false
		}
	}
	return true
}
