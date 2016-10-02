package main

import (
	"fmt"
	"sync"
	"time"
)

// START1 OMIT
func main() {
	err := hoge()
	fmt.Println(err)
	time.Sleep(time.Second * 5)
	fmt.Println("done")
}

func consumer(errc chan<- error) {
	errc <- fmt.Errorf("error") // writer // HL
	fmt.Println("error occurred")
}

var concurrency = 3 // ３並列で動かす // HL

// END1 OMIT

// START2 OMIT
func hoge() error {
	var wg sync.WaitGroup
	errc := make(chan error)
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
		case e := <-errc: // reader // HL
			return e
		case <-done:
			return nil
		default:
		}
	}
}

// END2 OMIT
