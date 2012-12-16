package main

import (
  "log"
  "github.com/mrb/sylvester"
)

func main() {
  log.Print("starting")

  graph := sylvester.NewGraph()

  input := graph.NewNode()
  process := graph.NewNode()
  output := graph.NewNode()

  graph.NewEdge(input, process)
  graph.NewEdge(process, output)

  graph.Activate()

  log.Print("finished")
}
