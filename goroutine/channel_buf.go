/*
同步通道-使用带缓冲的通道
一个无缓冲通道只能包含1个元素，有时显得很局限。我们给通道提供一个缓存，可以在扩展的make命令中
设置它的容量，如下：
buf := 100
ch1 := make(chan string, buf)
buf是通道可以同时容纳的元素（这里是string）个数
在缓冲满载之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，知道缓冲空了。
缓冲容量和类型无关，所以可以（尽量可能导致危险）给一些通道设置不同的容量，只要他们拥有同样的元素
类型。内置的cap函数可以返回缓冲区的容量。
如果容量大于0， 通道就是异步的了：缓冲满载（发送）或变空（接受）之前通信不会阻塞，元素会按照发送
的顺序被接受。如果容量是0或者未设置，通信仅在收发双方准备好的情况下菜可以成功。
ch := make(chan type, value)
value == 0 ->synchronous,unbuffered(阻塞)
value >0 -> asynchronous, buffered（非阻塞）取决于value元素
若使用通道的缓冲，你的程序会在请求激增的时候表现更好：更具弹性，专业术语叫：更具有伸缩性（scalable）
要在首要位置使用无缓冲通道设计算法，只在不确定的情况下使用缓冲。

协程中用通道输出结果
为了知道计算何时完成，可以通过信道回报。
ch := make(chan int)
go sum(bigArray, ch)
sum := <- ch
也可以使用通道来达到同步的目的，这个很有效的用法在传统计算机中称为信号量（semaphore）。或者换个方式，
通过通道发送信号告知处理已经完成（在协程中）。
在其他协程运行时让main()程序无线阻塞的通常做法是在main函数的最后放置一个select{}。
也可以使用通道让main程序等待协程完成，就是所谓的信号量模式

信号量模式
func compute(ch chan int) {
	ch <- someComputation()
}

func main() {
	ch := make(chan int)
	go compute(ch)
	doSomethingElseForWhile()
	result := <- ch
}

这个信号也可以是其他的，不返回结果
ch := make(chan int)
go func() {
	ch <- 1 
}()
doSomethingElseForAWhile()
<- ch

或者等待两个协程完成，每一个都会对切片S的一部分进行排序，片段
done := make(chan bool)
doSort := func(s []int){
	sort(s)
	done <- true
}
i:= privot(s)
go doSort(s[:i])
go doSort(s[i:])
<- done
<- done

下边的代码，用完整的信号量模式对长度为N的float64切片进行了N个doSomething()计算并同时完成，
通道sem分配了相同的长度（且包含空接口类型的元素），待所有的计算都完成后，发送信号（通过
放入值）。在循环中从通道sem不停地接受数据来等待所有的协程完成。
type Empty interface{}
var empty Empty
...
data := make([]float64, N)
res := make([]float64, N)
sem := make(chan Empty, N)
...
for i, xi := range data {
	go func (i int, xi float64) {
		res[i] = doSomething(i, xi)
		sem <- empty
	} (i, xi)
}
for i := 0; i < N; i++ { <- sem }
*/