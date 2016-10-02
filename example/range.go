package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	in := make(chan int, 5)
	go writer(in)
	reader(in)
	fmt.Println("fin")
}

func writer(in chan int) {
	for i := 0; i < 5; i++ {
		in <- i
		time.Sleep(time.Second * 1)
	}
	fmt.Println("finish writing")
}

func reader(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
	fmt.Println("in closed") // 5件処理したらすぐ抜ける？ // HL
}

// END OMIT
