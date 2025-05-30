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
	wg.Add(2)           // 等待两个协程完成
	chanA <- struct{}{} // 初始化，让A先执行
	go func() {         // 协程A：打印数字
		defer wg.Done()
		for i := 1; i < 6; i++ {
			<-chanA // 等待通道A的信号
			fmt.Printf("A:%d ", i)
			chanB <- struct{}{} // 向通道B发送信号
		}
	}()
	go func() { // 协程B：打印字母
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-chanB // 等待通道B的信号
			fmt.Printf("B:%c ", 'A'+i)
			chanA <- struct{}{} // 向通道A发送信号
		}
	}()
	wg.Wait()                      // 阻塞主线程，直到所有子协程完成
	fmt.Printf("\n%#v\n", <-chanA) // 此时chanA内还有数据
	close(chanA)
	close(chanB)
}

func TestAlternatePrintingUncacheV1(t *testing.T) {
	// 创建两个无缓存的通道
	chanA := make(chan struct{})
	chanB := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)   // 等待两个协程完成
	go func() { // 协程A：打印数字
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-chanA // 等待通道A的信号
			fmt.Printf("A:%d ", i+1)
			chanB <- struct{}{} // 向通道B发送信号

		}
	}()
	go func() { // 协程B：打印字母
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-chanB // 等待通道B的信号
			fmt.Printf("B:%c ", 'A'+i)
			if i != 4 { // 此时chanA已经不会再被读取了，程序会阻塞住这里
				chanA <- struct{}{} // 向通道A发送信号
			}
		}
	}()
	chanA <- struct{}{} // 初始化，让A先执行
	wg.Wait()           // 阻塞主线程，直到所有子协程完成
	close(chanA)
	close(chanB)
}

func TestAlternatePrintingUncacheV2(t *testing.T) {
	// 创建两个无缓存的通道
	chanA := make(chan struct{})
	chanB := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)   // 等待两个协程完成
	go func() { // 协程A：打印数字
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Printf("A:%d ", i+1)
			// 不进行初始化，直接由A开始发送，严格保证AB的发送和接收次数相同
			chanB <- struct{}{} // 向通道B发送信号
			<-chanA             // 等待通道A的信号
		}
	}()
	go func() { // 协程B：打印字母
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-chanB // 等待通道B的信号
			fmt.Printf("B:%c ", 'A'+i)
			chanA <- struct{}{} // 向通道A发送信号
		}
	}()
	wg.Wait() // 阻塞主线程，直到所有子协程完成
	close(chanA)
	close(chanB)
}
