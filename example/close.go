package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	done := make(chan struct{})
	go hoge(done)

	close(done)
	time.Sleep(time.Second * 2)
	fmt.Println("fin")
}

func hoge(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("done called")
			return // 必須 // HL
		}
	}
}

// END OMIT
