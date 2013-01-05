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

	node.NewSyncEvent(Sync)
	node.NewSyncEvent(Sync2)
	node.NewSyncEvent(Sync3)

	graph.Activate()

	<-graph.Control
}

func Starter(c syl.Channels) {
	cd := 0
	for {
		c.Data <- []byte{byte(cd)}
		cd += 1
		if cd > 100 {
			c.Control <- syl.NodeExit()
		}
	}
}

func Async(c syl.Channels) {
	for {
		select {
		case data := <-c.Data:
			<-time.After(time.Duration(rand.Int31n(100)) * time.Millisecond)
			log.Print("a", data)
		}
	}
}

func Sync(c syl.Channels) {
	select {
	case data := <-c.Data:
		<-time.After(time.Duration(rand.Int31n(100)) * time.Millisecond)
		log.Print("           s", data)
		c.Control <- syl.NodeNext()
	}
}

func Sync2(c syl.Channels) {
	select {
	case data := <-c.Data:
		<-time.After(time.Duration(rand.Int31n(100)) * time.Millisecond)
		log.Print("           s2", data)
		c.Control <- syl.NodeNext()
	}
}

func Sync3(c syl.Channels) {
	select {
	case data := <-c.Data:
		<-time.After(time.Duration(rand.Int31n(100)) * time.Millisecond)
		log.Print("           s3", data)
		c.Control <- syl.NodeSyncEventRestart()
	}
}

func Watcher(c syl.Channels) {
	for {
		<-time.After(10 * time.Millisecond)
		select {
		case control := <-c.Control:
			if bytes.Equal(control, syl.NodeExit()) {
				log.Print("Received EXIT, exiting")
				os.Exit(0)
			} else {
				c.Control <- control
			}
		}
	}
}
