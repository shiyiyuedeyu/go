/*
协程间的信道
概念：在第一个例子中，协程是独立执行的，它们之间没有通信。他们必须通信才会变得更有用：彼此之间发送和接受信息并且协调、同步他
们的工作。协程可以使用共享变量来通信，但是很不提倡这样做，因为这种方式给所有的共享内存的多线程都带来了苦难。
而go有一种特殊的类型，通道（channel），就像一个可以用于类型化数据的管道，由其负责协程之间的通信，从而避开所有由共享内存导致
的陷阱；这种通过通道进行通信的方式保证了同步性。数据在通道中进行传递：在任何给定时间，一个数据被设计为只有一个协程可以对其
访问，所以不会发生数据竞争。数据的所有权（可以读写数据的能力）也因此被传递。
工厂的传送带是个很有用的例子。一个机器（生产者协程）在传送带上放置物品，另外一个机器（消费者协程）拿到物品并打包。
通道服务于通信的两个目的：值的交换，同步的，保证了两个计算（协程）任何时候都是可知状态。
通常使用这样的格式来声明通道： var identifier chan datatype
未初始化的通道的值是nil
所以通道只能传输一种类型的数据，比如chan int或者chan string，所有的类型都可以用于通道，空接口interface{}也可以。甚至可以
（有时非常有用）创建通道的通道
通道实际上是类型化消息的队列：使得数据得以传输。它是先进先出FIFO的结构所以可以保证发送给他们的元素的顺序。通道也是引用类型，
所以我们使用make()函数来给它分配内存。这里先声明了一个字符串通道ch1，然后创建了它（实例化）：
var ch1 chan string
ch1 = make(chan string)
当然可以更短：ch1 := make(chan string)
这里我们构建一个int通道的通道：chanOfChans := make(chan int)
或者函数通道：funcChan := make(chan func())
所以通道是第一类对象：可以存储在变量中，作为函数的参数传递，从函数返回以及通过通道发送他们自身，另外它们是类型化的，
允许类型检查，比如尝试使用整数通道发送一个指针。

通信操作符 <-
这个操作符直观地标示了数据的传输：信息按照箭头的方向流动。
流向通道（发送）
ch <- int1
从通道流出（接收），三种方式：
int2 = <-ch 表示，变量int2从ch接收数据（获取新值）；假设int2已经声明过了，如果没有：int2 := <-ch
<- ch可以单独调用获取通道的（下一个）值，当前值会被丢弃，但是可以用来验证
if <- ch != 1000 {
	...
}
同一个操作符 <- 既可用于发送也可以接收，但Go会根据操作对象弄明白
通道的发送和接收都是原子操作：它们总是互不干扰的完成的。
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washingtong"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	for {
		input = <- ch
		fmt.Printf("%s ", input)
	}
}

/*
mian()函数中启动了两个协程：sendData()通过通道ch发送了5个字符串，getData()按顺序接收它们并打印出来。
如果两个协程需要通信，你必须给它们同一个通道作为参数
我们发现协程之间的同步非常重要：
main()等待了1秒让两个协程完成，如果不这样，sendData()就没机会输出。
getData()使用了无限循环：它随着sendData()的发送完成和ch变空也结束了
如果我们移除go，程序无法运行
为什么会这样？运行时（runtime）会检查所有的协程是否在等待着什么东西（可从某个通道读取或者写入某个通道），
这意味着程序将无法继续执行。这是死锁（deadlock）的一种形式，而运行时（runtime）可以为我们检测到这种情况。

*/