package workerpool

type Worker[E any] interface {
	Seed(buf chan<- E)
	Job(data E)
}
