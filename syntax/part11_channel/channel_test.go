package part11_channel

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	//var ch chan int
	//ch <- 123
	//val := <-ch
	//ch1 := make(chan int)
	ch2 := make(chan int, 2)
	ch2 <- 123
	val := <-ch2
	fmt.Println(val)
}

func TestChannelBlock(t *testing.T) {
	ch2 := make(chan int, 1)
	go func() {
		// 3秒后写入数据
		time.Sleep(time.Millisecond * 3000)
		ch2 <- 123

	}()
	fmt.Println(runtime.NumGoroutine())
	// 此处会阻塞三秒
	val := <-ch2
	fmt.Println(val)
}
