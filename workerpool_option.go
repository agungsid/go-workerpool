package workerpool

type WorkerPoolOption[T any] func(*WorkerPool[T])

func OptionID[T any](id string) WorkerPoolOption[T] {
	return func(wp *WorkerPool[T]) {
		wp.id = id
	}
}

func OptionNumWorker[T any](numWorker int) WorkerPoolOption[T] {
	return func(wp *WorkerPool[T]) {
		wp.numWorker = numWorker
	}
}

func OptionBufSize[T any](bufSize int) WorkerPoolOption[T] {
	return func(wp *WorkerPool[T]) {
		wp.bufSize = bufSize
	}
}

func OptionOnPanic[T any](onPanic func(issue interface{})) WorkerPoolOption[T] {
	return func(wp *WorkerPool[T]) {
		wp.onPanic = onPanic
	}
}
