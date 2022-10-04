package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	ch := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Handler(ch, i)
	}

	wg.Wait()
}

func Handler(ch chan int, id int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("This is routine # ", id)
		time.Sleep(time.Second * 1)
	}
}

func mainSecond() {

	ch := make(chan string)
	quit := make(chan string)

	wg.Add(1)
	go HandlerSelect(ch, quit)

	time.Sleep(time.Second * 1)
	ch <- "Hello Hamza!"

	time.Sleep(time.Second * 5)
	ch <- "Hello!"

	time.Sleep(time.Second * 2)
	quit <- "Bye!"
	wg.Wait()
	close(quit)
	close(ch)
}

func HandlerSelect(ch chan string, quit chan string) {
	defer wg.Done()
	for {
		time.Sleep(time.Second * 1)
		select {
		case val := <-ch:
			fmt.Println(val)
		case val := <-quit:
			fmt.Println(val)
			return
		default:
			fmt.Println("no action...")
		}
	}
}
