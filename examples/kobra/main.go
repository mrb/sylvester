package main

import (
	"bytes"
	syl "github.com/mrb/sylvester"
	"log"
)

func main() {
	log.Print("starting")

	graph := syl.NewGraph()

	input := graph.NewNode()

	doubler := func(id syl.DChan) (od syl.DChan, e syl.EChan) {
		nb := bytes.Map(func(r rune) rune {
			return rune(2 * int(r))
		}, <-id)

		od <- nb
		e <- nil

		return od, e
	}

	process := graph.NewNode()

	_ = process.AddEFunc(doubler)
	_ = process.AddEFunc(doubler)

	output := graph.NewNode()

	graph.NewEdge(input, process)
	graph.NewEdge(process, output)

	graph.Activate()

	log.Print("finished")
}
