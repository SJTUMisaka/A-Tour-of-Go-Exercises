package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for i := range ch1 {
		j, ok := <-ch2
		if !ok || i != j {
			return false
		}
	}
	_, ok := <-ch2
	return !ok
}

func testWalk() {
	ch := make(chan int)
	t := tree.New(1)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	//testWalk()
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println(Same(t1, t2))
}
