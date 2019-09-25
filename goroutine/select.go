/*
使用select切换协程
在不同的并发执行的协程中获取值可以通过关键字select来完成，他和switch控制语句
非常相似也被称作通信开关；它的行为像是轮询机制；select监听进入通道的数据，
也可以是用通道发送值的时候
select {
case u := <- ch1:
	...
case v := <- ch2:
	...
	...
default:
	...
}
default语句是可选的。在任何一个case中执行break或者return，select就结束了。
select做的就是：选择处理列出的多个通信情况中的一个
如果都阻塞了，会等待直到其中一个可以处理
如果多个可以处理，随机选择一个
如果没有通道操作可以处理并且写了default语句，它就会执行：default永远是可运行的
（这就是准备好了，可以执行）
在select中使用发送操作并且有的default可以确保发送不被阻塞！如果没有default，select
就会一直阻塞。
select语句实现了一种监听模式，通常用在（无线循环）中；在某种情况下，通过break语句使
循环退出。
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <- ch1 :
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <- ch2 :
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}