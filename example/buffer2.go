package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	in := writer()
	go reader(in)
	time.Sleep(time.Second * 3)
	close(in)
	fmt.Println("called close")
	time.Sleep(time.Second * 5)
}

func writer() chan int {
	elm := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	in := make(chan int, len(elm))
	for _, v := range elm {
		in <- v
	}
	return in
}

func reader(in chan int) {
	for i := range in {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

// END OMIT
