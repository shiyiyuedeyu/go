package main
import "fmt"

func main() {
	//缓冲区大小为2
	ch := make(chan int, 2)

	//因为ch是带缓冲的通道，我们可以同时发送两个数据
	//而不用立刻去同步读取数据
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}