package main

import (
	"fmt"
	"time"
)

func main() {
	hoge()
	time.Sleep(time.Second * 3)
}

// START OMIT
func hoge() {
	in := writer(1, 2, 3, 4, 5)

	go reader(in)

	return
}

func writer(nums ...int) <-chan int {
	in := make(chan int, len(nums))
	go func() {
		for _, n := range nums {
			in <- n
		}
		close(in)
	}()
	return in // HL
}

func reader(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

// END OMIT
