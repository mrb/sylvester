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
	byteSender := func(dc syl.DChan, ec syl.EChan) {
		dc <- []byte{1, 2, 3, 4, 5}
	}
	_ = input.AddEFunc(byteSender)

	process := graph.NewNode()
	doubler := func(dc syl.DChan, ec syl.EChan) {
		dc <- bytes.Map(func(r rune) rune {
			return rune(2 * int(r))
		}, <-dc)
	}
	_ = process.AddEFunc(doubler)

	processTwo := graph.NewNode()
	plusTenner := func(dc syl.DChan, ec syl.EChan) {
		dc <- bytes.Map(func(r rune) rune {
			return rune(10 + int(r))
		}, <-dc)
	}
	_ = processTwo.AddEFunc(plusTenner)

	output := graph.NewNode()
	stdOutPrinter := func(dc syl.DChan, ec syl.EChan) {
		log.Print("output --> ", <-dc)
		graph.ExitChan <- true
	}
	_ = output.AddEFunc(stdOutPrinter)

	// input -> process -> processTwp -> output
	graph.NewEdge(processTwo, output)
	graph.NewEdge(process, processTwo)
	graph.NewEdge(input, process)

	graph.Activate()

	<-graph.ExitChan

	log.Print("finished")
}
