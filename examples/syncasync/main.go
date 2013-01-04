package main

import (
	syl "github.com/mrb/sylvester"
	"log"
	"math/rand"
	"time"
)

func main() {
	graph := syl.NewGraph()

	node := graph.NewNode()

	node.NewAsyncEvent(Starter)
	node.NewAsyncEvent(Async)
	node.NewAsyncEvent(Async)
	node.NewAsyncEvent(Async)
	node.NewAsyncEvent(Async)

	node.NewSyncEvent(Sync)
	node.NewSyncEvent(Sync)
	node.NewSyncEvent(Sync)
	node.NewSyncEvent(Sync)

	graph.Activate()

	select {
	case <-graph.Control:
		log.Print("Peace")
	}
}

func Starter(c syl.Channels) {
	cd := 0
	for {
		<-time.After(time.Duration(rand.Int31n(150)) * time.Millisecond)
		c.Data <- []byte{byte(cd)}
		cd += 1
	}
}

func Async(c syl.Channels) {
	for {
		select {
		case data := <-c.Data:
			log.Print("a", data)
		}
	}
}

func Sync(c syl.Channels) {
	log.Print("here sync")

	for {
		select {
		case data := <-c.Data:
			log.Print("s", data)
		}
	}
}
