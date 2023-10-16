package main

import (
	"fmt"
	"sync"
	"time"
)

var st struct{ A, B, C int }

var mutex *sync.Mutex

func UpdateAndPrint(n int) {

	mutex.Lock()

	st.A = n
	time.Sleep(time.Millisecond)
	st.B = n
	time.Sleep(time.Millisecond)
	st.C = n
	time.Sleep(time.Millisecond)
	fmt.Println(st)

	mutex.Unlock()
}

func main() {

	mutex = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				UpdateAndPrint(j)
			}
		}()
	}

	// ゴルーチンが走る前にメインルーチンが終わらないようにしている
	for {

	}

}
