package workerpool

type Worker interface {
	Seed(buf chan<- interface{})
	Job(buf <-chan interface{})
}
