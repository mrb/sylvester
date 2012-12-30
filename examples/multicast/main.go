package main

import (
	syl "github.com/mrb/sylvester"
	"log"
	"time"
)

func main() {
	graph := syl.NewGraph()

	dataStream := graph.NewNode()
	dataStreamFunc := func(c syl.Channels) {
		ticker := time.NewTicker(time.Millisecond * 100)
		for _ = range ticker.C {
			log.Print("tick")
			c.Control.Ping()
		}

		log.Print("here?")

		select {
		case <-c.Control:
			log.Print("got it!")
			ticker.Stop()
		}
	}
	_ = dataStream.NewEvent(dataStreamFunc)

	outputOne := graph.NewNode()
	stdOutFunc := func(c syl.Channels) {
		for {
			select {
			case control := <-c.Control:
				log.Print("[1]", control, outputOne.Id())
			}
		}
	}
	_ = outputOne.NewEvent(stdOutFunc)

	outputTwo := graph.NewNode()
	stdOutFunc2 := func(c syl.Channels) {
		d := 0
		for {
			select {
			case control := <-c.Control:
				d++
				if d > 4 {
					c.Control <- syl.Exit()
					log.Print("blockin'")
				} else {
					log.Print("[2]", control, outputTwo.Id())
				}
			}
		}
	}
	_ = outputTwo.NewEvent(stdOutFunc2)

	graph.NewEdges(dataStream, []*syl.Node{outputOne, outputTwo})
	graph.NewEdge(outputTwo, dataStream)

	graph.Activate()

	<-graph.Control
	log.Print("Received Exit Signal, exiting")
}
