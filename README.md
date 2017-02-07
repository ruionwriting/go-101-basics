# go-101-basics
Introduction to Go, the tooling and basic language features

## Authors
Rob Reid [@codingconcepts](https://github.com/codingconcepts)

Nick Lanng [@nicklanng](https://github.com/nicklanng)

## Topics
1. Install Go
2. What is the Go workspace and why is it different to project workspaces
3. Data types
4. Slices and arrays
5. Pointers vs values
6. Functions
7. [FizzBuzz Exercise](#fizzbuzz-exercise)
8. String Calculator Exercise

<a name="fizzbuzz-exercise"/>
## 7. FizzBuzz Exercise

In this classic coding kata, you are tasked with printing the numbers from 1 to 100.
However, there are some extra rules.
If the number is a multiple of 3 then, instead of the number, print 'Fizz'.
If the number is a multiple of 5 then, instead of the number, print 'Buzz'.
If the number is a multiple of BOTH then instead of the number, print 'FizzBuzz'.

<details><summary>Click here to see the solution!</summary><p>
```go
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
```
</p></details>

