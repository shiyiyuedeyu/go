package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")
	go longWait()
	//longWait()
	go shortWait()
	//shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) //sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9)//sleep for 2 seconds
	fmt.Println("End of shortWait()")
}

/*
协程更有用的一个例子应该是在一个非常长的数组中查找一个元素
将数组分割为若干个不重复的切片，然后给每一个切片启动一个协程进行查找计算，
这样许多并行的协程可以用来进行查找任务，整体的查找时间会缩短（除以协程的数量）
*/
