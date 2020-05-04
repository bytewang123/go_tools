package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//case1()
	//case2()
	//case3()
	//case4()
	//case5()
	//case6()
	case7()
}

func case1() {
	lock := &sync.Mutex{}

	//谁先获得了互斥锁，必须等该协程解锁以后，另一个协程才能获得互斥锁
	go func() {
		fmt.Println("g1 start to get lock")
		lock.Lock()
		fmt.Println("g1 start to do sth")
		defer func() {
			fmt.Println("g1 start to unlock")
			time.Sleep(3 * time.Second)
			lock.Unlock()
			fmt.Println("g1 unlock")
		}()
	}()

	go func() {
		fmt.Println("g2 start to get lock")
		lock.Lock()
		fmt.Println("g2 start to do sth")
		defer func() {
			fmt.Println("g2 start to unlock")
			time.Sleep(3 * time.Second)
			lock.Unlock()
			fmt.Println("g2 unlock")
		}()
	}()

	time.Sleep(100 * time.Second)
}

func case2() {
	lock := &sync.RWMutex{}
	//多个协程可以同时获得读锁
	go func() {
		fmt.Println("g1 start to get rlock")
		lock.RLock()
		defer func() {
			fmt.Println("g1 start to unlock rlock")
			time.Sleep(5 * time.Second)
			lock.RUnlock()
			fmt.Println("g1 unlock rlock")
		}()
		fmt.Println("g1 do sth")
	}()

	go func() {
		fmt.Println("g2 start to get rlock")
		lock.RLock()
		defer func() {
			fmt.Println("g2 start to unlock rlock")
			time.Sleep(5 * time.Second)
			lock.RUnlock()
			fmt.Println("g2 unlock rlock")
		}()
		fmt.Println("g2 do sth")
	}()

	go func() {
		fmt.Println("g3 start to get rlock")
		lock.RLock()
		defer func() {
			fmt.Println("g3 start to unlock rlock")
			time.Sleep(5 * time.Second)
			lock.RUnlock()
			fmt.Println("g3 unlock rlock")
		}()
		fmt.Println("g3 do sth")
	}()

	time.Sleep(100 * time.Second)
}

func case3() {
	lock := &sync.RWMutex{}
	//多个协程可以同时获得读锁
	go func() {
		fmt.Println("g1 start to get rlock")
		lock.RLock()
		defer func() {
			fmt.Println("g1 start to unlock rlock")
			time.Sleep(5 * time.Second)
			lock.RUnlock()
			fmt.Println("g1 unlock rlock")
		}()
		fmt.Println("g1 do sth")
	}()

	go func() {
		fmt.Println("g2 start to get rlock")
		lock.RLock()
		defer func() {
			fmt.Println("g2 start to unlock rlock")
			time.Sleep(5 * time.Second)
			lock.RUnlock()
			fmt.Println("g2 unlock rlock")
		}()
		fmt.Println("g2 do sth")
	}()

	//写锁必须在所有读锁都已经释放后才能获取到
	go func() {
		fmt.Println("g3 start to get lock")
		lock.Lock()
		defer func() {
			fmt.Println("g3 start to unlock lock")
			time.Sleep(5 * time.Second)
			lock.Unlock()
			fmt.Println("g3 unlock lock")
		}()
		fmt.Println("g3 do sth")
	}()

	time.Sleep(100 * time.Second)
}

func case4() {
	lock := &sync.RWMutex{}

	//若先获取了写锁，其他协程无法获取读锁
	//直到写锁释放了，其他协程才能获取读锁
	fmt.Println("g3 start to get lock")
	lock.Lock()
	fmt.Println("g3 do sth")

	go func() {
		fmt.Println("g1 start to get rlock")
		lock.RLock()
		defer func() {
			fmt.Println("g1 start to unlock rlock")
			time.Sleep(5 * time.Second)
			lock.RUnlock()
			fmt.Println("g1 unlock rlock")
		}()
		fmt.Println("g1 do sth")
	}()

	go func() {
		fmt.Println("g2 start to get rlock")
		lock.RLock()
		defer func() {
			fmt.Println("g2 start to unlock rlock")
			time.Sleep(5 * time.Second)
			lock.RUnlock()
			fmt.Println("g2 unlock rlock")
		}()
		fmt.Println("g2 do sth")
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("g3 start to unlock lock")
	lock.Unlock()
	fmt.Println("g3 unlock lock")

	time.Sleep(100 * time.Second)
}

func case5() {
	once := &sync.Once{}
	for i := 0; i < 10; i++ {
		fmt.Printf("idx:%+v\n", i)
		//只有i=0会被执行
		once.Do(func() {
			fmt.Printf("do %+v\n", i)
		})
	}
}

func case6() {
	once := &sync.Once{}
	for i := 0; i < 10; i++ {
		idx := i
		go func() {
			fmt.Printf("idx:%+v\n", idx)
			//只有第一次执行到once的idx会被执行
			once.Do(func() {
				fmt.Printf("do %+v\n", idx)
			})
		}()
	}
	time.Sleep(5 * time.Second)
}

func case7() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go do(i)
	}
	time.Sleep(5 * time.Second)
}

func do(i int) {
	fmt.Printf("idx:%+v\n", i)
}
