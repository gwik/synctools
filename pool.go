// Additions to go sync package. it provides a Pool structure that limit
// the concurrency of goroutines.
package synctools

import "sync"

// A Pool spawns jobs (goroutines) and limit the concurrency.
type Pool struct {
	wg         *sync.WaitGroup
	completion chan bool
}

// Builds a new Pool limiting concurrency by concurrencyLimit
func NewPool(concurrencyLimit int) *Pool {
	wg := sync.WaitGroup{}
	completionChan := make(chan bool, concurrencyLimit)
	for i := 0; i < concurrencyLimit; i++ {
		completionChan <- true
	}
	return &Pool{&wg, completionChan}
}

// Spawns job in a new goroutine, waiting if concurrency limit is reached.
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

// Wait for the completion of all the spawned jobs.
func (pool *Pool) Wait() {
	pool.wg.Wait()
}
