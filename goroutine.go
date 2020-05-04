package main

import (
	"fmt"
	"time"
)

func main() {
	//case1()
	//case2()
	//case3()
	//case4()
	//case5()
	//case6()
	//case7()
	//case8()
	//case9()
	//case10()
	//case11()
	case12()
}

func case1() {
	c := make(chan int, 1)
	go push(c)

	for i := range c {
		fmt.Printf("get %+v\n", i)
	}
}

func push(c chan int) {
	for {
		c <- 1
		fmt.Printf("push %+v\n", 1)
		time.Sleep(1 * time.Second)
	}
}

func case2() {
	c := make(chan int, 1)
	go func() {
		for i := range c {
			fmt.Printf("get %+v\n", i)
		}
	}()
	push(c)
}

//写了两次，读了三次，读最后一次时管道已经空了
//且没有别的goroutine存在，所以死锁
func case3() {
	c := make(chan int, 1)
	go func() {
		c <- 1
		c <- 2
	}()
	//这里为了测试增加了一个测试协程
	//这个协程什么也没做，但是一直在sleep
	//这时整个程序并不会死锁
	//直到sleep时间结束才会出现死锁
	//所以判断死锁的标志应该是当一个协程尝试从一个空channel取数据
	//但是[没有别的协程存在]
	go func() {
		time.Sleep(100 * time.Second)
		fmt.Println("test...")
	}()

	fmt.Printf("get %+v\n", <-c)
	fmt.Printf("get %+v\n", <-c)
	fmt.Println("extra...")
	fmt.Printf("get %+v\n", <-c)
}

//case1, case2, 有协程在不断地读管道，若读不到了，管道为空，但是发送协程是一个无限循环
//不会退出，因此对于读取协程来说，总是还有别的协程存在
//因此不会死锁

func case6() {
	c := make(chan int, 1)
	c <- 1
	//管道的容量为1，
	//但是有两次写操作，没有读操作
	//第二次写操作会报死锁
	c <- 2
	fmt.Println("finish")
}

func case7() {
	c := make(chan int, 1)
	c <- 1
	fmt.Println("write 1 finish")

	//写完第一次后，开一个测试协程
	//这个协程什么也不做
	//程序不会报死锁，能正常走到write 2 start
	//当sleep结束，程序报死锁
	//因此可以看下面总结
	go func() {
		fmt.Println("start to test...")
		time.Sleep(100 * time.Second)
		fmt.Println("test...")
	}()

	fmt.Println("write 2 start")
	c <- 2
	fmt.Println("write 2 finish")
}

func case8() {
	c := make(chan int, 1)
	go func() {
		//容量为1，但是写了三次
		//不会死锁
		//因为向管道写数据，管道已满，发现还有别的协程存在（主协程）
		//因此不会死锁
		c <- 1
		c <- 2
		c <- 3
		c <- 4
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func case9() {
	c := make(chan int, 1)
	//读协程不会死锁
	//虽然一开始是在向一个空管道读数据，但是还有主协程存在，不会死锁
	go func() {
		fmt.Println(<-c)
		fmt.Println(<-c)
	}()

	//下面的操作会死锁
	//因为读协程读完两个数据就退出了
	//管道的容量是1，因此3可以正常写入，这时管道刚刚满
	//写4的时候就不行了，因为管道已经装不下了
	//死锁发生
	c <- 1
	c <- 2
	c <- 3
	c <- 4
}

//总结：
//当向一个管道写数据，管道已满，且无别的协程存在，死锁发生
//当从一个管道读数据，管道已空，且无别的协程存在，死锁发生

func case11() {
	c := make(chan int, 1)
	go func() {
		//容量为1，但是写了三次
		//不会死锁
		//因为向管道写数据，管道已满，发现还有别的协程存在（主协程）
		//因此不会死锁
		fmt.Printf("send:%+v\n", 1)
		c <- 1
		fmt.Printf("send:%+v\n", 2)
		c <- 2
		fmt.Printf("send:%+v\n", 3)
		c <- 3
		fmt.Printf("send:%+v\n", 4)
		c <- 4
		fmt.Println("close")
	}()

	//如果发送端不close管道，以下遍历会发生死锁
	//这是因为，当遍历到最后4之后，再去读管道，
	//此时是向一个空管道读数据，
	//而写管道的协程已经关闭了，没有别的协程了
	//因此死锁发生
	for i := range c {
		fmt.Printf("get:%+v\n", i)
	}
}

func case12() {
	c := make(chan int, 1)
	go func() {
		//容量为1，但是写了三次
		//不会死锁
		//因为向管道写数据，管道已满，发现还有别的协程存在（主协程）
		//因此不会死锁
		fmt.Printf("send:%+v\n", 1)
		c <- 1
		fmt.Printf("send:%+v\n", 2)
		c <- 2
		fmt.Printf("send:%+v\n", 3)
		c <- 3
		fmt.Printf("send:%+v\n", 4)
		c <- 4
		fmt.Println("close")
	}()
	go func() {
		fmt.Println("start to test...")
		time.Sleep(100 * time.Second)
		fmt.Println("test...")
	}()
	//发送端不close管道，以下遍历会在test协程结束后发生死锁
	//与case11的区别是多了一个测试协程
	//即遍历c，取到4之后，向一个空管道取数据，本应死锁
	//但是测试协程还存在，因此不会死锁
	//当测试协程结束，死锁发生
	for i := range c {
		fmt.Printf("get:%+v\n", i)
	}
}

//无缓冲管道和有缓冲管道本质上是相同的
//只是无缓冲管道一来就是满的
func case4() {
	ch := make(chan int)
	fmt.Println("send...")
	//无缓冲管道一来就是满的，
	//再往里发数据，就会导致死锁
	ch <- 1
	fmt.Println("finish send...")
	fmt.Println(<-ch)
}

func case5() {
	ch := make(chan int)
	fmt.Println("send...")
	//无缓冲管道一来就是满的，
	//再往里发数据，就会导致死锁
	ch <- 1
	fmt.Println("finish send...")
	go func() {
		fmt.Println(<-ch)
	}()
}

//不会死锁
func case10() {
	ch := make(chan int)
	//从一个空管道读数据，读不到
	//但是还有主协程存在，因此不会死锁，暂时阻塞等待
	go func() {
		fmt.Println(<-ch)
	}()
	fmt.Println("send...")
	//往一个满的管道写数据，会死锁
	//但是此时，读取协程正在阻塞等待仍然存在，
	//因此不会死锁
	//与case5的区别，是将写管道的位置挪到了读取协程的后面
	//这样读取管道协程先被创建
	//主协程的写管道操作本应死锁，但是发现有读取协程
	//因此不会死锁了
	ch <- 1
	fmt.Println("finish send...")
}
