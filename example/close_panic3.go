package main

import "fmt"

func main() {
	ch := make(chan int)
	// START OMIT
	close(ch)
	h := <-ch // zero値が生成される
	fmt.Println(h)
	// END OMIT
}
