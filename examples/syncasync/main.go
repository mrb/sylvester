/*

syncasync - a simple program to show sync and async Events cohabiting in the
            same node.

*/

package main

import (
	"bytes"
	syl "github.com/mrb/sylvester"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// A program needs at least one graph
	graph := syl.NewGraph()

	// A graph needs at least one node
	node := graph.NewNode()

	// A node needs some Events! That's where stuff happens.
	// This one's simple - it streams bytes of data. It's below.
	node.NewAsyncEvent(dataStreamer)

	// This Event is important - we pass it the node and the graph, and we
	// expect it to do important work.
	node.NewAsyncEvent(makeWatcher(node, graph))

	// You can add a lot of events to a node. They interact with each other in
	// different ways. 3 "asyncLogger" Event functions are instantiated here.
	node.NewAsyncEvents(makeAsyncLogger("a0"), makeAsyncLogger("a1"), makeAsyncLogger("a2"))

	// These three Events are different - they're "sync". They'll get executed
	// in order and their scheduling is controlled by an external function. In 
	// this case, that's the watcher Event.
	node.NewSyncEvents(makeSyncLogger("s0"), makeSyncLogger("s1"), makeSyncLogger("s2"))

	// "Activate" is Sylvester's word for "start the flow of data."
	graph.Activate()

	// Block on the graph's Control channel receive - a simple naive mechanism
	// to ensure everything else stays running. We'll send it a signal from the
	// watcher function (that's the reason we pass the grap into its closure) and
	// then the app will wait for cleanup before exiting.
	<-graph.Control
	log.Print("Received EXIT, exiting in 100ms")
	<-time.After(100 * time.Millisecond)
	os.Exit(0)
}

// Loop forever, streaming bytes.
func dataStreamer(c syl.Channels) {
	for {
		for cd := 0; ; cd++ {
			c.Data <- []byte{byte(cd)}
		}
	}
	c.Control.Exit()
}

// This function gets a pointer to the Node and the Graph passed in
// to its closure. This is up to the user's discretion. Control signals are
// handled by the watcher, including scheduling sync events and handling Exit
// signals.
func makeWatcher(node *syl.Node, graph *syl.Graph) syl.Event {
	return func(c syl.Channels) {
		for {
			select {
			case control := <-c.Control:
				switch {
				case bytes.Equal(control, syl.NodeExit()):
					graph.Control.Exit()
				case bytes.Equal(control, syl.NodeNext()):
					node.NextSyncEvent()
				default:
					log.Print("Unhandled control event: ", control)
				}
			}
		}
	}
}

// Returns a function that will loop forever, pull data off the data
// channel, and print it.
func makeAsyncLogger(name string) syl.Event {
	return func(c syl.Channels) {
		for {
			select {
			case data := <-c.Data:
				<-time.After(time.Duration(rand.Int31n(10)) * time.Millisecond)
				log.Print(name, data)
			}
		}
	}
}

// Returns a function that will run once, blocking until it receives data,
// and then printing it out. When it's done, it signals the Control channel
// to proceed to the next sync Event.
func makeSyncLogger(name string) syl.Event {
	return func(c syl.Channels) {
		select {
		case data := <-c.Data:
			<-time.After(time.Duration(rand.Int31n(10)) * time.Millisecond)
			log.Print("       ", name, data)
			c.Control.Next()
		}
	}
}
