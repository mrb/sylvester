package sylvester

import (
	"log"
)

type Edge struct {
	id    []byte
	anode *Node
	bnode *Node
}

func (e *Edge) Activate() {
	go func() {
		for {
			select {
			case data := <-e.anode.output:
				log.Print(data)
				e.bnode.input <- data
			}
		}
	}()
}
