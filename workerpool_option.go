package workerpool

type WorkerPoolOption func(*WorkerPool)

func OptionID(id string) WorkerPoolOption {
	return func(wp *WorkerPool) {
		wp.id = id
	}
}

func OptionNumWorker(numWorker int) WorkerPoolOption {
	return func(wp *WorkerPool) {
		wp.numWorker = numWorker
	}
}

func OptionBufSize(bufSize int) WorkerPoolOption {
	return func(wp *WorkerPool) {
		wp.bufSize = bufSize
	}
}

func OptionOnPanic(onPanic func(issue interface{})) WorkerPoolOption {
	return func(wp *WorkerPool) {
		wp.onPanic = onPanic
	}
}
