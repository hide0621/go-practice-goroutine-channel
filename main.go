package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := new(sync.WaitGroup)

	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Groutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Groutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Groutine")
		}
		wg.Done()
	}()

	wg.Wait()

}
