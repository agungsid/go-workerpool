# go-workerpool
Worker-Pool written in GO

## Installation
```bash
go get github.com/agungsid/go-workerpool
```

## Usage
```go
package main

type SampleSeeder struct{}

func (s *SampleSeeder) Seed(buff chan<- interface{}) {
	for i := 0; i < 1000; i++ {
		buff <- i
	}
}

func (s *SampleSeeder) Job(data interface{}) {
	i, _ := data.(int)
	log.Println(i)
}

func main() {
	pool := workerpool.NewWorkerPool(&SampleSeeder{}, workerpool.OptionID("sample-workerpool"))
	pool.Do()
}
```
