package main

import "fmt"

// START OMIT
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch) // 先にclose // HL

	for n := range ch {
		fmt.Println(n) // 表示は後
	}
	fmt.Println("done")
}

// END OMIT
