package main

func main() {
	ch := make(chan int)
	// START OMIT
	close(ch)
	close(ch) // panic
	// END OMIT
}
