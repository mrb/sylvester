package main

import (
	syl "github.com/mrb/sylvester"
	"log"
	"time"
)

func main() {
	graph := syl.NewGraph()

	dataStream := graph.NewNode()
	dataStreamFunc := func(dc syl.DataChan, ec syl.ErrorChan) {
		ticker := time.NewTicker(time.Millisecond * 500)
		for _ = range ticker.C {
			log.Print("tick")
			dc <- []byte{9, 9, 9, 9}
		}
	}
	_ = dataStream.NewEvent(dataStreamFunc)

	outputOne := graph.NewNode()
	stdOutFunc := func(dc syl.DataChan, ec syl.ErrorChan) {
		for {
			select {
			case data := <-dc:
				log.Print("[1]", data, outputOne.Id())
			}
		}
	}
	_ = outputOne.NewEvent(stdOutFunc)

	outputTwo := graph.NewNode()
	stdOutFunc2 := func(dc syl.DataChan, ec syl.ErrorChan) {
		for {
			select {
			case data := <-dc:
				log.Print("[2]", data, outputTwo.Id())
			}
		}
	}
	_ = outputTwo.NewEvent(stdOutFunc2)

	graph.NewEdge(dataStream, []*syl.Node{outputOne, outputTwo})

	graph.Activate()

	<-graph.ExitChan
	log.Print("Received Exit Signal, exiting")
}
