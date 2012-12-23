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
		ec <- nil
	}
	_ = input.AddEFunc(byteSender)

	process := graph.NewNode()
	doubler := func(dc syl.DChan, ec syl.EChan) {
		select {
		case data := <-dc:
			newData := bytes.Map(func(r rune) rune {
				return rune(2 * int(r))
			}, data)

			dc <- newData
			ec <- nil
		}
	}
	_ = process.AddEFunc(doubler)

	output := graph.NewNode()
	stdOutPrinter := func(dc syl.DChan, ec syl.EChan) {
		select {
		case data := <-dc:
			log.Print("output --> ", data)
			dc <- data
			ec <- nil
			graph.ExitChan <- true
		}
	}
	_ = output.AddEFunc(stdOutPrinter)

	graph.NewEdge(process, output)
	graph.NewEdge(input, process)

	graph.Activate()

	<-graph.ExitChan

	log.Print("finished")
}
