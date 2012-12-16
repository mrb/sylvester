package main

import (
  "github.com/mrb/sylvester"
)

func main() {
  log.Print("starting")

  graph := NewGraph()

  input := graph.NewNode()
  process := graph.NewNode()
  output := graph.NewNode()

  graph.NewEdge(input, process)
  graph.NewEdge(process, output)

  graph.Activate()

  a := func([]byte) []byte {
    return nil
  }

  log.Print(reflect.TypeOf(a))

  log.Print("finished")
}
