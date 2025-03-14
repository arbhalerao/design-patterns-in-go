package main

import (
	"sync"

	"github.com/abha1erao/design-patterns-in-go/creational/singleton"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			singleton.GetInstance()
			// fmt.Printf("%p\n", instance)
		}()
	}

	wg.Wait()
}
