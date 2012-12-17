package sylvester

import (
	"log"
)

type Edge struct {
	id    []byte
	anode *Node
	bnode *Node
}

func NewEdge(anode, bnode *Node) *Edge {
	return &Edge{
		id:    newID(),
		anode: anode,
		bnode: bnode,
	}
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

func (e *Edge) Id() *[]byte {
	return &e.id
}
