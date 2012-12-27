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
	worker := graph.NewNode()
	output := graph.NewNode()

	graph.NewEdge(supervisor, worker)
	graph.NewEdge(worker, output)
	graph.Activate()

	os.Exit(0)
}
