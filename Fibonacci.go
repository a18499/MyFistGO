package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func getINT() int {
	return 8;
}
func main() {
	const tset  = 15
	b :=5
	f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f())
	fmt.Println("This is my first IDEA GO")
	fmt.Println(getINT()+b)
}

