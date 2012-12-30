package main

import (
	syl "github.com/mrb/sylvester"
	"log"
	"os"
)

func main() {
	log.Print("New Graph")
	graph := syl.NewGraph()

	supervisor := graph.NewNode()
	superFunc := func(c syl.Channels) {
		log.Print("Supa")
		for {
			select {
			case <-c.Control:
				log.Print("Got control signal")
			}
		}
	}
	_ = supervisor.NewEvent(superFunc)

	worker := graph.NewNode()
	workerFunc := func(c syl.Channels) {
		log.Print("Worka")

		for {
			select {
			case <-c.Control:
				log.Print("Got control signal")
			}
		}
	}
	_ = worker.NewEvent(workerFunc)

	output := graph.NewNode()
	stdOutFunc := func(c syl.Channels) {
		log.Print("Printa")

		for {
			select {
			case control := <-c.Control:
				log.Print("[O]", control)
			}
		}
	}
	_ = output.NewEvent(stdOutFunc)

	graph.NewEdge(supervisor, worker)
	graph.NewEdge(worker, output)

	graph.Activate()

	<-graph.Data

	os.Exit(0)
}
