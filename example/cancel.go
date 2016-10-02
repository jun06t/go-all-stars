package main

import (
	"fmt"
	"time"
)

func main() {
	err := hoge()
	fmt.Println(err)
	time.Sleep(time.Second * 5)
	fmt.Println("done")
}

// START1 OMIT
func hoge() error {
	done := make(chan struct{})
	defer close(done)
	in := writer(done)
	go reader(in)
}

func reader(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

// END1 OMIT

// START2 OMIT
func writer(done <-chan struct{}) <-chan int {
	nums := []int{1, 2, 3, 4, 5}
	in := make(chan int, len(nums))
	go func() {
		defer close(in)
		for _, n := range nums {
			time.Sleep(time.Second * 1)

			select { // HL
			case <-done: // doneが来たら終了する // HL
				return // HL
			case in <- n:
			default:
			}
		}
	}()
	return in
}

// END2 OMIT
