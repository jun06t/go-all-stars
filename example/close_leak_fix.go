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
	fmt.Println("done")
}

func hoge() error {
	done := make(chan struct{})       // HL
	defer close(done)                 // 追加 // HL
	in := writer(done, 1, 2, 3, 4, 5) // HL
	go reader(in)

	time.Sleep(time.Second * 3)
	return fmt.Errorf("error occurred")
}

// END1 OMIT

// START2 OMIT
func writer(done <-chan struct{}, nums ...int) <-chan int {
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

func reader(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

// END2 OMIT
