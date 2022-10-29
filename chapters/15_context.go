package main

import (
	"context"
	"fmt"
	"time"
)

// Golang Context is a tool that is used to share request-scoped data,
// cancellation signals, and timeouts or deadlines across API layers or
// processes in a program. It is one of the most important tools while
// working with concurrent programming in Go.

func main() {
	// Contexts are always created with other contexts. But there are two ways of creating
	// a context from scratch AKA root contexts. Root contexts are created with Background or TODO methods
	ctx := context.Background()
	_ = context.TODO() // use this as placeholder context for when you dont know what kind of context you want to use.

	// A context is just an interface that looks like this defined in context package
	type Context interface {
		Deadline() (deadline time.Time, ok bool)
		Done() <-chan struct{}
		Err() error
		Value(key interface{}) interface{}
	}

	// These are empty context objects out of which we can create a child contexts to pass around from function to function
	// or go routine to go routine. The use of these child contexts is to pass data between scopes and goroutines.
	// The data can be actual data, or a cancel signal that tells any functions with these contexts to perform necessary actions

	// For example; you can also cancel network requests...
	// when a server get a request you can grab the context that comes bundled with the request obj and listen on whether
	// its been cancelled if the request was cancelled by client. We can use this to cancel or prevent any costly computations, IO, db etc...

	// There are quite a few ways to create a child context. We can create another child context from a child context. We can create
	// context trees so that if one context is cancelled all its children receive cancellation signal.

	// As we know, Root contexts are created with Background or TODO methods,
	// while derived contexts are created using WithCancel, WithDeadline, WithTimeout, or WithValue methods.
	valctx := context.WithValue(ctx, "auth_token", "XYZ_123")

	fmt.Println(valctx.Value("auth_token"))

	cancelctx, cancelfunc := context.WithCancel(ctx)

	// cancelctx has a channel called done which receives an empty struct and closes itself when cancelfunc is called.
	// We can listen on the done channel and stop all execution if anything is receuved using select statements
	// its always a good idea to defer cancelfunc later just in case we dont manually end up calling cancel func
	defer cancelfunc()

	timeoutctx, cancelfunc := context.WithTimeout(ctx, 3*time.Second)

	// it cancels when cancelfunc is called or timeout exceeds whatever comes first

	deadlinectx, cancelfunc := context.WithDeadline(ctx, time.Now().Add(5*time.Second))

	// it cancels when cancelfunc is called or deadline reached whatever comes first...

	fmt.Println(cancelctx, timeoutctx, deadlinectx)
}
