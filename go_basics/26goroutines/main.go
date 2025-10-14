package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

// wg => wait group

var signals = []string{"test"}

var wg sync.WaitGroup  // usually these are pointer

// Mutex => A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex. A mutex must not be copied after first lock.
// (lock or unlock mutex) What mutex is basically does => i am going to lock this memeory till this one goroutine is working till the time it is writing anything inside that i will not allow anybody to just use this memeory.

var mut sync.Mutex // pointer

func main() {
  // go greeter("Hello") // the goroutine is simply created by adding a keyword go
	// greeter("World")

	websitelist := []string{
		"https://go.dev",
		"https://meta.com",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
		"https://x.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait() // Always goes at the end of main method and whatever the method is been called out
	// wg.Wait() hey main please don't get exit yet some of my friends are coming in and the job of wg.Add() is keep on adding the goroutines number how many of my friends have been out there. in this case we have just 1 because just 1 call happening but their can a sutation their might be 5 or 10 different others.

	fmt.Println(signals)

}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string){
	defer wg.Done() 

  res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
    mut.Unlock()
		
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

