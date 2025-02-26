package main

import (
	"context"
	"log"
	"math/rand/v2"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	/*
		Create a semaphore of a capacity of size 2
		This means that go routines will be processed atmost 2 in concurently
		since the semaphore size is at two
	*/
	/*sem := semaphore.NewWeighted(2)

	// Simulating 100 concurrent requests
	for i := 0; i < 100; i++ {
		go func(id int) {
			// Acquire the semaphore
			err := sem.Acquire(context.Background(), 1)
			if err != nil {
				log.Printf("Request %d could not be processed, semaphore full: %v", id, err)
				return
			}
			defer sem.Release(1)

			// Simulate a request taking 1 second to process
			time.Sleep(time.Second)
			log.Printf("Request %d processed successfully", id)
		}(i)
	}

	// Wait for all requests to complete
	time.Sleep(time.Second * 5)
	log.Println("All requests processed")*/
	pool := semaphore.NewWeighted(2)
	var wg sync.WaitGroup
	wg.Add(6)
	go swim("Candier", pool, &wg)
	go swim("polite", pool, &wg)
	go swim("Cooler", pool, &wg)
	go swim("greener", pool, &wg)
	go swim("pinker", pool, &wg)
	go swim("Party", pool, &wg)
	wg.Wait()
	log.Println("Main: Done, shutting down")
}

func swim(name string, pool *semaphore.Weighted, wg *sync.WaitGroup) {
	defer pool.Release(1)
	defer wg.Done()
	log.Printf("%v: I want to swim and I can get a lane\n", name)
	// In real applications, pass in your context such as HTTP request context
	ctx := context.Background()
	// We can also Acquire/Release more than 1
	// when the workloads consume different amount of resources
	if err := pool.Acquire(ctx, 1); err != nil {
		log.Printf("%v: Ops, something went wrong! I cannot acquire a lane\n", name)
		return
	}

	durationSwim := time.Duration(rand.IntN(5))
	log.Printf("%v: I was able to get/release a lane(resource), I will swim for %d seconds\n", name, durationSwim)
	time.Sleep(durationSwim * time.Second)
	log.Printf("%v: I finished swimming and Releasing the lane(resource)\n", name)
}
