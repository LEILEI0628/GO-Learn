package part11_channel

import (
	"context"
	"fmt"
	"runtime"
	"sync"
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
	fmt.Println(runtime.NumGoroutine()) // goroutine数量
	// 此处会阻塞三秒
	val := <-ch2
	fmt.Println(val)
}

func SafeClose(ch chan int) {
	_, ok := <-ch
	if ok {
		close(ch)
	}
}

// 这个 ch 一定是 MyStruct 来关
type MyStruct struct {
	ch        chan struct{}
	ctx       context.Context
	closeOnce sync.Once
}

// 用户会多次调用，或者多个 goroutine 调用
func (m *MyStruct) Close() error {
	m.closeOnce.Do(func() {
		// 确保整个代码只会执行一次
		close(m.ch)
	})
	return nil
}

//func (m *MyStruct) UseV1(ch chan struct{}) error {
//	UseV2(ch)
//	UseV3(ch)
//}
//
//type ChUsage struct {
//	ch chan int
//}

type MyStructBV1 struct {
	// 暴露出去了，你就不知道用户啥时候会给你关了
	Ch        chan struct{}
	closeOnce sync.Once
}

func TestLoopChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 100)
		}
		close(ch)
	}()
	for val := range ch {
		t.Log(val)
	}

	t.Log("channel 被关了")

	//for {
	//	val, ok := <-ch
	//	if !ok {
	//		break
	//	}
	//	t.Log(val)
	//}
}

func TestSelect(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 123
	ch2 <- 234
	//go func() {
	//	time.Sleep(time.Millisecond * 100)
	//	ch1 <- 123
	//}()
	//go func() {
	//	time.Sleep(time.Millisecond * 100)
	//	ch2 <- 123
	//}()
	select {
	case val := <-ch1:
		t.Log("ch1", val)
		val = <-ch2
		t.Log("ch2", val)
	case val := <-ch2:
		t.Log("ch2", val)
		val = <-ch1
		t.Log("ch1", val)
	}

	println("往后执行")
}

func TestGoroutineCh(t *testing.T) {
	ch := make(chan int)
	// 这一个就泄露掉了
	go func() {
		// 永久阻塞在这里
		ch <- 123
	}()

	// 这里后面没人往 ch 里面读数据
}

func TestGoroutineChRead(t *testing.T) {
	ch := make(chan int, 100000)
	// 这一个就泄露掉了
	def := new(BigObj)
	go func() {

		// 永久阻塞在这里
		for i := 0; i < 100000; i++ {
			ch <- i
		}
		abc := new(BigObj)
		t.Log(abc)
		t.Log(def)
		// 永久阻塞在这里，ch 占据的内存，永远不会被回收
		ch <- 1
	}()

	// 这里后面没人往 ch 里面读数据
}

type BigObj struct {
}

func ReturnReadWriteChannel() chan struct{} { // 返回可读写channel
	panic("implement me")
}

func ReturnReadOnlyChannel() <-chan struct{} { // 返回只读channel
	panic("implement me")
}

func ReturnWriteOnlyChannel() chan<- struct{} { // 返回只写channel
	// 很少使用，可以直接要求返回一个channel
	panic("implement me")
}
