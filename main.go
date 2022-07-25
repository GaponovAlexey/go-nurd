package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	r := time.Now()
	count := 0
	wiw(count)
	fmt.Println("time:", time.Since(r).Microseconds())

}

func wiw(count int) {
	var wa sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wa.Add(1)
		go func() {
			count++
		defer wa.Done()
		}()
		wa.Wait()
	}
	
	fmt.Println(count)
	fmt.Println("exit")
}
