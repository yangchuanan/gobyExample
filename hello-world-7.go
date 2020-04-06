package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func goroutineDemo() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("gong")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func chanDemo1() {
	// 创建一个新的通道
	message := make(chan string)
	go func() {
		// 向通道中发送一个值
		message <- "ping"
	}()

	// 从通道中接收一个值
	msg := <-message
	fmt.Println(msg)
}

func chanDemo2() {
	// 通道缓冲
	message := make(chan string, 2)
	message <- "buffered"
	message <- "channel"
	fmt.Println(<-message)
	fmt.Println(<-message)
}

func worker(done chan bool) {
	fmt.Println("woking...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

//通道同步
func chanSyncDemo() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}

// 只写通道
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 可读可写通道
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func chanDirectionDemo() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func chanSelectDemo() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 同时等待c1 和c2
	for i := 0; i < 2; i++ {
		//通道选择器
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

}

func timeoutDemo() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
