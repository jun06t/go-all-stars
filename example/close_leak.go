package main

import (
	"fmt"
	"time"
)

// START1 OMIT
func main() {
	err := hoge()
	fmt.Println(err)
	time.Sleep(time.Second * 5)
	fmt.Println("fin")
}

func hoge() error {
	in := writer(1, 2, 3, 4, 5)
	go reader(in)

	time.Sleep(time.Second * 3)
	return fmt.Errorf("error occurred") // エラーの発生 // HL
}

// END1 OMIT

// START2 OMIT
func writer(nums ...int) <-chan int {
	in := make(chan int, len(nums))
	go func() {
		defer close(in)
		for _, n := range nums {
			time.Sleep(time.Second * 1)
			in <- n
		}
	}()
	return in
}

func reader(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

// END2 OMIT
