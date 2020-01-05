package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"math/rand"
	"time"
)

type routinePool struct {
	pool *tunny.Pool
}

type task struct {
	data int
}

type executor struct {
	pool *routinePool
}

func main() {

	e := &executor{}
	p := &routinePool{}
	p.pool = tunny.NewFunc(3, e.worker) //注册调度者的worker方法
	e.pool = p
	//defer e.pool.Close()
	for i := 0; i < 10; i++ {
		e.push(i) //调度者把任务推送
	}
	time.Sleep(200 * time.Second)
}

func (e *executor) push(v int) {
	go e.pool.process(v) //go协程开始做
}

func (p *routinePool) process(v int) {
	task := &task{data: v}
	p.pool.Process(task) //最终协程池的Process方法会调度到注册的worker方法中
}

func (e *executor) worker(work interface{}) interface{} {
	switch w := work.(type) {
	case *task:
		err := e.doTask(w.data) //调度者的worker方法，实际执行真实任务
		return err
	}
	return fmt.Errorf("task type error")
}

//实际执行真实任务的方法
func (e *executor) doTask(v int) error {
	fmt.Printf("start do %+v...\n", v)
	fmt.Printf("this is = %+v\n", v)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("%+v done...\n", v)
	return nil
}
