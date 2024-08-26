//go:build ignore 

package main

import (
	"fmt"
	"sync"
	"time"
)

// in many programming languages with concurrency features - they rely on OS threads
// to implement concurrency. As we know OS threads are heavy processes with thread memory
// exceeding 1MB and the OS scheduler is resposible for coordinating many other threads as well.
// Context switching is also a big problem here. So we need to start introducing entities like thread pools and etc.

// Because of all of this sometimes the gains of concurrency is overshadowed by the
// operational overhead of the OS threads

// In go, we have an abstraction over an OS thread with golang providing its own
// thread implementation thats lightweight and its own scheduler with more efficient
// scheduling algorithm.

// This abstraction is simply called go routine

// To create a go routine we need to simply use the `go` keyword and
// pass a function invocation to it.

// It has all the usual problems with using concurrency - i.e, thread waiting, race conditions,
// deadlocks and starvation.

// These synchronization problems can be solved by using waitgroups or mutexes from the sync package

func printTheMsg() {
	fmt.Println("This is from a go routine")
}

func noWaitContainer() {
	go printTheMsg() // This will not be printing anything
	// The main funciton is also a go routine - this ends befire the new printTheMsg routine
	// gets a chance to finish executing
}

func waitContainer() {
	go printTheMsg()
	// in order for the main routine to finish until the spawned go routine
	// to finish executing we need to make sure the main routine is running long enough
	// but that doesnt mean we can stick a bunch of code and hope the spawned go routine
	// has enough time to finish executing

	// we have to make sure that the main routine has paused running its work in a non blocking way
	// so that the scheduler can say ok time for next routine to execute

	// An infinite loop for example is a blocking code
	// an io read is a non blocking work
	// a sleep is a non blocking work

	// lets sleep here so that the spawned routine gets a chance to start
	// before the main routine exits
	time.Sleep(5 * time.Millisecond) // 5 seconds
	// When main routine reached above line it stops and lets the next available routine start execution.
}

func closureContainer() {
	msg := "hello from main routine"
	go func() {
		fmt.Println(msg) // the msg variable is here is coming from a closure
	}()

	// go routines can access variables in its closure!

	time.Sleep(5 * time.Millisecond) // 5 seconds
}

func closureWithRaceContainer() {
	msg := "hello from main routine"
	go func() {
		fmt.Println(msg) // the msg variable is here is coming from a closure
		// However by the time msg is printed the msg variable is reassigned to "Goodbye"...
	}()
	msg = "Goodbye from main routine"

	// go routines can access variables in its closure!

	time.Sleep(5 * time.Millisecond) // 5 seconds
}

func closureWithRaceContainerSolved() {
	msg := "hello from main routine"
	go func(msg string) {
		fmt.Println(msg) // To solve the race condition from above just pass
		// the data dependencies as an argument to the routine
		// the argument is evaluated before golang stacks this routine to be executed.
	}(msg)
	msg = "Goodbye from main routine"

	// go routines can access variables in its closure!

	time.Sleep(5 * time.Millisecond) // 5 seconds
}

func waitGroupsContainer() {
	// using a sleeper wait is obviously not a production worthy code.
	// however we still need to wait for the routine to get a chgance to execute
	// before the main routine is finished.

	// We can use a wait group for this. Its kinda like `thread.join` in ruby
	// WaitGroup is a struct in sync package - structs are used to mimic classes in golang
	var waitGroup sync.WaitGroup
	waitGroup = sync.WaitGroup{}

	// wait group is basically a counter. Increment this counter and launching a go routine
	// and decrement this counter when the go routine is done. In the main routine done exit until the counter is zero.

	waitGroup.Add(2)
	go func() {
		fmt.Println("wait group container routine 1")
		waitGroup.Done() // Decrements the counter by 1
	}()
	go func() {
		fmt.Println("wait group container routine 2")
		waitGroup.Done() // Decrements the counter by 1
	}()

	waitGroup.Wait() // This is basically a non-blocking wait loop until counter = 0
}

// Another way to synchronize threads is to do mutexes
func withoutMutexContainer() {
	counter := 0
	waitGroup := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		go func() {
			fmt.Println(counter)
			waitGroup.Done()
		}()
		go func() {
			counter++
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}

func withMutexContainer() {
	mutex := sync.RWMutex{}
	// This is a read write mutex: From documentation
	// A RWMutex is a reader/writer mutual exclusion lock. The lock can be held by
	// an arbitrary number of readers or a single writer. The zero value for a
	// RWMutex is an unlocked mutex.
	counter := 0
	waitGroup := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		go func() {
			mutex.RLock()
			fmt.Println(counter)
			mutex.RUnlock()
			waitGroup.Done()
		}()
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

func main() {
	fmt.Println("go routine patterns in golang - uncomment only one at a time to run\n")

	// noWaitContainer() // doesnt print anything
	// waitContainer() // This is from a go routine
	// closureContainer() // hello from main routine
	// closureWithRaceContainer() // Goodbye from main routine
	// closureWithRaceContainerSolved() // hello from main routine
	// waitGroupsContainer() // "wait group container routine 1" and "wait group container routine 2" in a non-deterministic order
	// withoutMutexContainer() // Youre gonna a non-deterministic random order of numbers here everytime you run as
	// withMutexContainer() // you get deterministic 1 through 10 everytime you run this
}
