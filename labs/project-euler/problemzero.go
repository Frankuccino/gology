package main

import (
	"fmt"
)

func main() {
	var sum int = 0
	for i := range 609000 {
		if i%2 != 0 {
			square := int(i) * int(i)
			sum += square
		}
	}

	fmt.Println(sum)
}

// Problem statement:
// A number is a perfect square, or a square number, if it is the square of a positive integer.
// For example, 25 is a square number because  5^2   ; it is also an odd square.
//
// The first 5 square numbers are: 1, 4, 9, 16, 25, and the sum of the odd squares is    .
//
// Among the first 609 thousand square numbers, what is the sum of all the odd squares?
// Answer for the sum of the square roots: 92774763510
// Answer for Sum of the Squares: 37644421499898500
