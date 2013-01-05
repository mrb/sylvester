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
	node.NewAsyncEvent(Watcher)

	node.NewAsyncEvent(ASyncLogger)
	node.NewAsyncEvent(ASyncLogger)
	node.NewAsyncEvent(ASyncLogger)

	node.NewSyncEvent(SyncLogger)
	node.NewSyncEvent(SyncLogger2)
	node.NewSyncEvent(SyncLogger3)

	graph.Activate()

	<-graph.Control
}

func Starter(c syl.Channels) {
	for cd := 0; cd < 100; cd++ {
		c.Data <- []byte{byte(cd)}
	}
	c.Control <- syl.NodeExit()
}

func ASyncLogger(c syl.Channels) {
	for {
		select {
		case data := <-c.Data:
			<-time.After(time.Duration(rand.Int31n(10)) * time.Millisecond)
			log.Print("a", data)
		}
	}
}

func SyncLogger(c syl.Channels) {
	select {
	case data := <-c.Data:
		<-time.After(time.Duration(rand.Int31n(10)) * time.Millisecond)
		log.Print("           s", data)
		c.Control <- syl.NodeNext()
	}
}

func SyncLogger2(c syl.Channels) {
	select {
	case data := <-c.Data:
		<-time.After(time.Duration(rand.Int31n(10)) * time.Millisecond)
		log.Print("           s2", data)
		c.Control <- syl.NodeNext()
	}
}

func SyncLogger3(c syl.Channels) {
	select {
	case data := <-c.Data:
		<-time.After(time.Duration(rand.Int31n(10)) * time.Millisecond)
		log.Print("           s3", data)
		c.Control <- syl.NodeSyncEventRestart()
	}
}

func Watcher(c syl.Channels) {
	for {
		select {
		case control := <-c.Control:
			if bytes.Equal(control, syl.NodeExit()) {
				log.Print("Received EXIT, exiting in 100ms")
				<-time.After(500 * time.Millisecond)
				os.Exit(0)
			} else {
				c.Control <- control
				<-time.After(1 * time.Millisecond)
			}
		}
	}
}
