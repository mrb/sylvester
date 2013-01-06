/*

An example illustrating the use of async and sync events within the same node.
The "Starter" event pushes bits into the data channel, and the async logger is
a simple consume-and-log event. The sync events each consume a bit and log, and
signal when they are done to trigger the next sync event. The results of this
example will show that the async events are returned in the order dictated by
the random amount of time they sleep, while the sync events always execute and
print bits in a guaranteed order.

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
	graph := syl.NewGraph()

	node := graph.NewNode()

	node.NewAsyncEvent(Starter)

	node.NewAsyncEvent(Watcher(node))

	node.NewAsyncEvent(ASyncLogger)
	node.NewAsyncEvent(ASyncLogger)
	node.NewAsyncEvent(ASyncLogger)

	node.NewSyncEvent(SyncLogger)
	node.NewSyncEvent(SyncLogger2)
	node.NewSyncEvent(SyncLogger3)

	graph.Activate()

	<-graph.Control
	log.Print("Received EXIT, exiting in 100ms")
	<-time.After(100 * time.Millisecond)
	os.Exit(0)
}

func Starter(c syl.Channels, g syl.ControlChan) {
	for {
		for cd := 0; cd < 100; cd++ {
			c.Data <- []byte{byte(cd)}
		}
	}
	c.Control.Exit()
}

func Watcher(node *syl.Node) syl.Event {
	return func(c syl.Channels, g syl.ControlChan) {
		for {
			select {
			case control := <-c.Control:
				switch {
				case bytes.Equal(control, syl.NodeExit()):
					g.Exit()
				case bytes.Equal(control, syl.NodeNext()):
					node.NextSyncEvent()
				default:
					log.Print("Unhandled control event: ", control)
				}
			}
		}
	}
}

func ASyncLogger(c syl.Channels, g syl.ControlChan) {
	for {
		select {
		case data := <-c.Data:
			<-time.After(time.Duration(rand.Int31n(10)) * time.Millisecond)
			log.Print("a", data)
		}
	}
}

func SyncLogger(c syl.Channels, g syl.ControlChan) {
	select {
	case data := <-c.Data:
		<-time.After(time.Duration(rand.Int31n(10)) * time.Millisecond)
		log.Print("           s", data)
		c.Control.Next()
	}
}

func SyncLogger2(c syl.Channels, g syl.ControlChan) {
	select {
	case data := <-c.Data:
		<-time.After(time.Duration(rand.Int31n(10)) * time.Millisecond)
		log.Print("           s2", data)
		c.Control.Next()
	}
}

func SyncLogger3(c syl.Channels, g syl.ControlChan) {
	select {
	case data := <-c.Data:
		<-time.After(time.Duration(rand.Int31n(10)) * time.Millisecond)
		log.Print("           s3", data)
		c.Control.Next()
	}
}
