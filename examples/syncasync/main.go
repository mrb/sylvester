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

	node.NewAsyncEvent(func(c syl.Channels, g syl.ControlChan) {
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
	})

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
	for cd := 0; cd < 100; cd++ {
		c.Data <- []byte{byte(cd)}
	}
	c.Control.Exit()
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
