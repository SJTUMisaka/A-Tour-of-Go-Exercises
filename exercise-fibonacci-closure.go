package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	firstnum := -1
	secondnum := 1
	return func() int {
		fibnum := firstnum + secondnum
		firstnum = secondnum
		secondnum = fibnum
		return fibnum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
