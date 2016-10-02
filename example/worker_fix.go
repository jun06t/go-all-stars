package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	err := hoge()
	fmt.Println(err)
	time.Sleep(time.Second * 5)
	fmt.Println("done")
}

func consumer(errc chan<- error) {
	errc <- fmt.Errorf("error")
	fmt.Println("error occurred")
}

var concurrency = 3

// START OMIT
func hoge() error {
	var wg sync.WaitGroup
	errc := make(chan error, concurrency) // worker数分bufferを用意 // HL
	done := make(chan struct{})
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			consumer(errc)
		}()
	}
	go func() {
		wg.Wait()
		close(done)
	}()
	for {
		select {
		case e := <-errc:
			return e
		case <-done:
			return nil
		default:
		}
	}
}

// END OMIT
