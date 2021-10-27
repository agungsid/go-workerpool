package workerpool

import (
	"sync"

	"github.com/google/uuid"
)

type WorkerPool struct {
	id        string
	numWorker int
	bufSize   int
	onPanic   func(interface{})

	worker Worker

	buf chan interface{}

	running bool
	wg      sync.WaitGroup
}

func NewWorkerPool(worker Worker, options ...WorkerPoolOption) *WorkerPool {
	pool := &WorkerPool{
		id:     uuid.NewString(),
		worker: worker,
	}

	for _, opt := range options {
		opt(pool)
	}

	return pool
}

func (pool WorkerPool) ID() string {
	return pool.id
}

func (pool WorkerPool) NumWorker() int {
	numWorker := pool.numWorker
	if pool.numWorker <= 0 {
		numWorker = 10
	}
	return numWorker
}

func (pool WorkerPool) BufSize() int {
	bufSize := pool.bufSize
	if pool.bufSize <= 0 {
		bufSize = 10
	}
	return bufSize
}

func (pool *WorkerPool) recoverPanic() {
	r := recover()
	if r != nil {
		pool.onPanic(r)
	}
}

func (pool *WorkerPool) doSeed() {
	defer pool.recoverPanic()
	defer pool.wg.Done()
	pool.worker.Seed(pool.buf)
}

func (pool *WorkerPool) doJob() {
	defer pool.recoverPanic()
	defer pool.wg.Done()
	pool.worker.Job(pool.buf)
}

func (pool *WorkerPool) Do() {
	if pool.running {
		return
	}
	pool.running = true

	pool.wg.Add(pool.NumWorker() + 1)
	for i := 0; i < pool.NumWorker(); i++ {
		go pool.doJob()
	}

	go pool.doSeed()
	pool.wg.Wait()
	pool.running = false
}