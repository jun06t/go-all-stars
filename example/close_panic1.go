package main

func main() {
	ch := make(chan int)
	// START OMIT
	close(ch)
	ch <- 10 // panic
	// END OMIT
}
