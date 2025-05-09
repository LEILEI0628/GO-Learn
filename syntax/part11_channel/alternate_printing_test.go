package part11_channel

import (
	"fmt"
	"sync"
	"testing"
)

func TestAlternatePrinting(t *testing.T) {
	// 创建两个带缓存的通道（容量为1）
	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)   // 等待两个协程完成
	go func() { // 协程A：打印数字
		defer wg.Done()
		for i := 1; i < 6; i++ {
			<-chanA // 等待通道A的信号
			fmt.Printf("A: %d\n", i)
			chanB <- struct{}{} // 向通道B发送信号
		}
	}()
	go func() { // 协程B：打印字母
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-chanB // 等待通道B的信号
			fmt.Printf("B: %c\n", 'A'+i)
			chanA <- struct{}{} // 向通道A发送信号
		}
	}()
	chanA <- struct{}{} // 初始化，让A先执行
	wg.Wait()           // 阻塞主线程，直到所有子协程完成
}
