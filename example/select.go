package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(100 * time.Millisecond)
	var count int
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
			count++
		case <-boom:
			fmt.Println("BOOM!")
			count++
		}
		if count > 10 {
			return
		}
	}
}

// END OMIT
