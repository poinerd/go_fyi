package main

import(
    "fmt"
	"sync"
)

type Counter struct {
	mu sync.RWMutex
	count int
}

//I would use a wg instead of a channel

func main(){

	aWorker := &Counter{
		count : 0,
	}

	// var w sync.WaitGroup
	done := make(chan bool)
    
	// w.Add(2)

	// go work(aWorker, &w)
	// go work(aWorker, &w)

	go work(aWorker, done)
	go work(aWorker, done)

	// w.Wait()
	<-done
	<-done

	fmt.Printf("Final Counter Value: %d (Expected: 2000)\n", aWorker.count)


}

func work (c *Counter, done chan bool){
	// defer w.Done()
	c.mu.Lock()
	defer c.mu.Unlock()

	for i:=0 ; i<1000 ; i++{
		c.count ++
	}

	done<-true
}






