package main

import (
	"fmt"
	"sync"

	"github.com/agungsid/go-workerpool"
)

type Sum struct {
	result int
	mu     sync.Mutex
}

func (c *Sum) Seed(buf chan<- int) {
	for i := 1; i <= 1000; i++ {
		buf <- i
	}
}

func (c *Sum) Job(data int) {
	c.mu.Lock()
	c.result += data
	c.mu.Unlock()
}

func main() {
	sum := &Sum{}
	wp := workerpool.NewWorkerPool[int](sum)
	fmt.Println("id:", wp.ID())
	fmt.Println("buf size:", wp.BufSize())
	fmt.Println("num worker:", wp.NumWorker())
	wp.Do()
	fmt.Println("result:", sum.result)
}
