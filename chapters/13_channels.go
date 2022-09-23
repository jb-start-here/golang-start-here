package main

import (
	"fmt"
	"time"
)

// Channels are a way for two go routines to talk to each other...
// you can declare a channel by using chan keyword followed by the the type of the messages
// that this channel will transport

func printChannels() {
	// a channel once defined is locked to a specific type of data transport only. its not data type agnostic

	var myChannel chan int         // declare a channel called myChannel that will only carry int type data
	var anotherChannel chan string // declare a channel called anotherChannel that will only carry string type data
	// the zero value of a channel is nil

	fmt.Println(myChannel)      // <nil>
	fmt.Println(anotherChannel) // <nil>

	// nil channels are not of any use and hence the channel has to
	// be defined using make similar to maps and slices.
	// myChannel = chan int wont work because it reassigning nil channel again
	// to a variable that already has nil value
	// as opposed to slices/maps we can assign a non nil but empty slice/map to a declared var

	// to assign
	myChannel = make(chan int)
	anotherChannel = make(chan string)

	// by default channels are bidirectional - a go routine can send and receive

	// to send message : `anotherChannel <- "Hello"`
	// ro receive message: `message := <- anotherChannel`

	// Let this go routine recieve the message
	go func() {
		msg := <-anotherChannel
		fmt.Println("The msg from anotherChannel is", msg)

		// lets send the message as well
		myChannel <- 45
	}()

	// Let this go routine send the message
	anotherChannel <- "Hello"
	//Lets receive from myChannel
	fmt.Println("The message from myChannel is", <-myChannel)

	// If we dont want to assign the msg to any variable just do `<- channel`
}

// Reading and sending messages from a channel is blocking operation
// This gives us the ability to synchronize go routines without locks and mutexes.

// All we need to do is to create a channel thats sends a done message!
// and send the done message from the go routine. The main go routine is then listenting
// on the channel from the done msg. It will not exit until a msg has been received

func syncRoutinesWithChannels() {
	doneChannel := make(chan bool)

	go func() {
		fmt.Println("Doing work in go routine....")
		time.Sleep(3 * time.Second) // sleep for 3 seconds to pretend to work
		fmt.Println("work finished in go routine....")
		doneChannel <- true
	}()

	// wait until the routine above is complete
	fmt.Println("waiting for done msg from done channel...now waiting")
	<-doneChannel
	fmt.Println("Received done msg from done channel...now exiting")
}

// One important factor to consider while using channels is deadlock.
// If a Goroutine is sending data on a channel, then it is expected that some
// other Goroutine should be receiving the data.
// If this does not happen, then the program will panic at runtime with Deadlock.
// Go runtimes are smart enough to detect that any routine is waiting for a data on
// a channel that will never occur so instead of being blocked and potentially starving other routines (if no parallelism)
// if will panic.

// its the same with goroutines that send msg to channel but no other routine
// is going to read it

func deadlockedRoutine() {
	x := make(chan int)

	<-x // deadlock fatal error
	// Exception is the we cant recover from deadlock panic
}

// We can also make unidirectional channels
// '<- chan' - send only
// 'chan <-' - receive only

// Whats the point of writing to a write only channel
// whats the point reading from a read only channel
// if no one is on the other end. It will cause deadlocks.

// Thats why we can case uni directional <-> bidirectional channels
// We do that while defining funcion args on go routines
func castToUnidirectional() {
	var bidirectional chan bool
	bidirectional = make(chan bool)

	go func(sendOnly chan<- bool) {
		sendOnly <- true
	}(bidirectional)

	<-bidirectional
}

// Senders have the ability to close the channel using `close(channel)` built in method to
// notify receivers that no more data will be sent on the channel.
// Everytime we receive data from a channel it actually give us two return values
// a val and ok

// ok is true if the value was received by a successful send operation to a channel.
// If ok is false it means that we are reading from a closed channel.
// The value read from a closed channel will be the zero value of the channel's type

// When a sender closes a channel one last message will be sent with (zeroVal, false) return type

func closeChannel() {
	channel := make(chan string)

	go func() {
		channel <- "Hello"
		close(channel)
	}()

	for {
		val, ok := <-channel
		if ok {
			fmt.Println("MSG:", val)
		} else {
			fmt.Println("Channel closed")
			break
		}
	}
}

// We can loop over channel using a range - it just the same as in closeChannel
// behind the scenes.

func rangeOnChannel() {
	channel := make(chan string)

	go func() {
		channel <- "Hello 1"
		channel <- "Hello 2"
		channel <- "Hello 3"
		close(channel)
	}()

	// Once channel is closed, the loop automatically exits.
	// Usually range gives us two returns the index and value but for ranging
	// over channels it just gives one return value
	for msg := range channel {
		fmt.Println("for MSG:", msg)
	}
}

// Buffered channels are channels with an inbuilt buffer to handle processing of
// msgs on both ends at a different rate

// unlike normal channels these are not blocking send and receives on a go routine
// To create a buffered channel just use make but pass in a second argument, an int to
// indicate buffer capacity
func bufferedChannels() {
	bChannel := make(chan string, 5)
	// it can hold 5 messages at max

	bChannel <- "Hello 1" // in a non buffered it would deadlock here
	bChannel <- "Hello 2"
	bChannel <- "Hello 3"
	bChannel <- "Hello 4"
	bChannel <- "Hello 5"

	// It's possible to read data from a already closed
	// buffered channel. The channel will return the data that is
	// already written to the channel and once all the data has been read,
	// it will return the zero value of the channel.
	close(bChannel)

	fmt.Println(<-bChannel)
	fmt.Println(<-bChannel)
	fmt.Println(<-bChannel)
	fmt.Println(<-bChannel)
	fmt.Println(<-bChannel)
	_, ok := <-bChannel
	fmt.Println(ok) // false
}

// We can use len and cap to figure out the max capacity of a channel
// the current length of the channel like slices

// The select statement is used to choose from multiple
// send/receive channel operations. The select statement blocks until
// one of the send/receive operations is ready. If multiple operations are ready,
// one of them is chosen at random. The syntax is similar to switch except that
// each of the case statements will be a channel operation.

func selectChannel() {
	output1 := make(chan string)
	output2 := make(chan string)
	go func(ch chan string) {
		time.Sleep(6 * time.Second)
		ch <- "from server1"
	}(output1)
	go func(ch chan string) {
		time.Sleep(3 * time.Second)
		ch <- "from server2"
	}(output2)

	// The fastest go routine will be selected
	// The other ignored
	// If we have a default case then that will be executed.

	// server 2 will always win
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

func main() {
	printChannels()
	syncRoutinesWithChannels()
	// deadlockedRoutine()
	castToUnidirectional()
	closeChannel()
	rangeOnChannel()
	bufferedChannels()
	selectChannel()
}
