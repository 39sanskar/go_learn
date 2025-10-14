package main

import (
	"fmt"
	"sync"
)


func main(){
	fmt.Println("Race condition - CoadingHub.in")

	wg := &sync.WaitGroup{}
  mut := &sync.RWMutex{}
 
	
	var score = []int{0}
  

	// func(){}() => function body and immediately execute.
	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("One Routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Two Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Three Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Reader Routine")
		mut.RLock()
		fmt.Println("Current score:", score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println("Final score:", score)
}

// their is no guarenty on order it might change or not change 
// Also Read/Write Mutex is exist

// go run --race .

/*

- A Read-Write Mutex (also known as RWMutex) is a synchronization primitive used in concurrent programming to protect shared resources while allowing multiple readers or a single writer, but not both at the same time.

- 🔹 Why we need Read-Write Mutex
- In normal sync.Mutex, only one goroutine can access a resource at a time — whether it’s reading or writing.
- However, in many cases, reads are safe to perform concurrently, and writes are rare.

That’s where sync.RWMutex helps:
- Multiple goroutines can read the shared data at the same time.
- When a goroutine writes, it gets exclusive access — no one else can read or write until it’s done.

*/