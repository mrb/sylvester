package main

import (
  "log"
  "github.com/mrb/sylvester"
)

func main() {
  log.Print("starting")

  graph := sylvester.NewGraph()

  input := graph.NewNode(sylvester.TCP)
  process := graph.NewNode(sylvester.Ionode)
  output := graph.NewNode(sylvester.TCP)

  graph.NewEdge(input, process)
  graph.NewEdge(process, output)

  graph.Activate()

  log.Print("finished")
}
