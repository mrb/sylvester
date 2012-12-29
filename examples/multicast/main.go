package main

import (
	syl "github.com/mrb/sylvester"
	"log"
	"time"
)

func main() {
	graph := syl.NewGraph()

	dataStream := graph.NewNode()
	dataStreamFunc := func(dc syl.DataChan, cc syl.ControlChan, ec syl.ErrorChan) {
		ticker := time.NewTicker(time.Millisecond * 10)
		for _ = range ticker.C {
			log.Print("tick")
			cc.Ping()
		}
	}
	_ = dataStream.NewEvent(dataStreamFunc)

	outputOne := graph.NewNode()
	stdOutFunc := func(dc syl.DataChan, cc syl.ControlChan, ec syl.ErrorChan) {
		for {
			select {
			case control := <-cc:
				log.Print("[1]", control, outputOne.Id())
			}
		}
	}
	_ = outputOne.NewEvent(stdOutFunc)

	outputTwo := graph.NewNode()
	stdOutFunc2 := func(dc syl.DataChan, cc syl.ControlChan, ec syl.ErrorChan) {
		for {
			select {
			case control := <-cc:
				log.Print("[2]", control, outputTwo.Id())
			}
		}
	}
	_ = outputTwo.NewEvent(stdOutFunc2)

	graph.NewEdges(dataStream, []*syl.Node{outputOne, outputTwo})

	graph.Activate()

	<-graph.Channels.Control
	log.Print("Received Exit Signal, exiting")
}
