package main

import (
	"fmt"
	"time"
)

func f(value string) {

	for i := 0; i < 3; i++ {
		fmt.Println(value)
		time.Sleep(3 * time.Second)
	}

}

func main() {

	// ゴルーチンの勉強

	// go f("goroutineを使って実行")

	// f("普通に実行")

	// fmt.Println("done")

	// 簡単なチャネルの勉強

	// messages := make(chan string)

	// go func() {
	// 	messages <- "Hello"
	// }()

	// msg := <-messages

	// fmt.Println(msg)

	// バッファとデッドロックの勉強

	ch := make(chan int, 2)

	// ch <- 1
	// ch <- 2
	// ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
