package part11_channel

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	//var ch chan int // 声明了一个存放int类型的channel
	//ch <- 123 // 声明但没有初始化，读写都会panic
	//val := <-ch
	// 空结构体一般用来做信号
	//var chs chan struct{}
	//ch1 := make(chan int) // 无容量
	ch2 := make(chan int, 2) // 固定容量，容量不会变化
	ch2 <- 123
	val := <-ch2
	fmt.Println(val)
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 123
	ch <- 456
	val, ok := <-ch
	t.Log(val, ok)
	close(ch)
	//ch <- 111 // 会panic：关闭后可以读但不能写

	val, ok = <-ch
	t.Log(val, ok)

	val, ok = <-ch // 在channel关闭的情况下会读到0 false，未关闭会等待写入（此处为阻塞）
	t.Log(val, ok)
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
