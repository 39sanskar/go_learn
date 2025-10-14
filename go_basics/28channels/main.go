package main

import (
	"fmt"
	"sync"
)

func main(){
  fmt.Println("Channels in golang")

	myCh := make(chan int)
  wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5  // push a 5 value into that, arrow always points towards left <- , never see this kind of arrow -> 
	
  
	// the channel works in this way, hey i am only allowing you to pass me a value if somebody is listening to me then only allow you to pass on me a value.

	// create goroutine 
	wg.Add(2)
	// Receive Only
	// channel is a box, so <-chan this one means their is a value which is going ontside of the box, receiving a value.
	go func(ch <-chan int, wg *sync.WaitGroup){
		
		val, isChanelOpen := <-myCh

    fmt.Println(isChanelOpen)
		fmt.Println(val)
    // fmt.Println(<-myCh) // listening the value of channel
		// fmt.Println(<-myCh)
		
		wg.Done()
	}(myCh, wg)

	// send only
	// inside the box send a value
	go func(ch chan<- int, wg *sync.WaitGroup){
		myCh <- 0
		close(myCh)
    
		// myCh <- 6
		
		wg.Done()
	}(myCh, wg)

	wg.Wait() // the end of main method  wg.Wait() 
}



/*

channels => channels are a way in which your multiple golang routine can actually talk to each other. they will still not be aware of what happening inside that what's long it take to another goroutine to finish up the job.

// may be you are waiting for just some signal or just some information from another go routine, you don't need to comeback and finish  all of that you can do it on the go while the thread has not completed the execution.

*/

// listening on the close channel we are receiving 0.

// when using  val, isChanelOpen := <-myCh  you are still receive a value 0 but based on this false and true decide the 0 is coming because of the close channel, or it is coming up because somebody who is trying to send a value.

// if receive false => 0 is coming because of the close channel
// if receive true => 0 is coming because somebody who is trying to send.