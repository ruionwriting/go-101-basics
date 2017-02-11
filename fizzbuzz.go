package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(FizzBuzz(i))
	}
}

// FizzBuzz takes an integer i and returns a string based
// on the following rules:
// - If the integer is divisible by 3 and 5 -> "FizzBuzz"
// - If the integer is divisible by 3       -> "Fizz"
// - If the integer is divisible by 5       -> "Buzz"
// - If the integer is divisible by neither -> i
func FizzBuzz(i int) string {
	switch {
	case i%15 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return "Fizz"
	case i%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(i)
	}
}
