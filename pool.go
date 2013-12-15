package syncext

import "sync"

type Pool struct {
	wg         *sync.WaitGroup
	completion chan bool
}

func NewPool(concurrencyLimit int) *Pool {
	wg := sync.WaitGroup{}
	completionChan := make(chan bool, concurrencyLimit)
	for i := 0; i < concurrencyLimit; i++ {
		completionChan <- true
	}
	return &Pool{&wg, completionChan}
}

func (pool *Pool) Spawn(job func()) {
	<-pool.completion
	pool.wg.Add(1)
	go func() {
		defer func() {
			pool.completion <- true
			pool.wg.Done()
		}()
		job()
	}()
}

func (pool *Pool) Wait() {
	pool.wg.Wait()
}
