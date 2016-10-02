package main

import (
	"fmt"
	"time"
)

func main() {
	hoge()
}

// START OMIT
func hoge() {
	in := make(chan int, 5)
	go writer(in)
	go reader(in)

	time.Sleep(time.Second * 3)
	fmt.Println("done")
}

func writer(in chan<- int) {
	for i := 0; i < 10; i++ {
		in <- i
	}
}

func reader(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

// END OMIT
